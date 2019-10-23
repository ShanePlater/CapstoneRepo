package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	auth "gopkg.in/korylprince/go-ad-auth.v2"
)

func getMe(g *gin.Context, m *models.Context) {
	var data types.GetMeJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Check if keyword is empty.
	if data.Username == "" || data.Password == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss username or password."})
		return
	}

	//Authenticate through AD before creating a backend session
	//Setup AD connection config
	config := &auth.Config{
		Server:   "LR-BNE-ERP-01",       /* Box that contains domain controller */
		Port:     389,                   /* Default Port for AD connection */
		BaseDN:   "DC=lar,DC=net,DC=au", /* Domain */
		Security: auth.SecurityNone,
	}

	//Connect to AD so we can get groups
	connectAD, err := config.Connect()
	if err != nil {
		fmt.Println("Error connecting to server:")
		fmt.Println(err)
		g.JSON(http.StatusUnauthorized, gin.H{"error": "Active Directory Connection failed"})
	}
	fmt.Println(connectAD.Conn)
	BindUPN := data.Username + "@lar.net.au"
	status, err := connectAD.Bind(data.Username, data.Password)
	if !status {
		fmt.Println("controllers/adauth.go Bind failed - username/pw error")
		g.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}
	if err != nil {
		fmt.Println("controllers/adauth.go error when authenticating")
		fmt.Println(err)
		return
	}

	//get user attributes
	entry, err := connectAD.GetAttributes("userPrincipalName", BindUPN, []string{"memberOf"})
	if err != nil {
		fmt.Println("error getting attributes from user")
		fmt.Println(err)
	}
	dnGroups := entry.GetAttributeValues("memberOf")
	fmt.Println(entry)
	approved := false

	//Loop through all of the users groups and check if they are in Redbook admins
	for _, group := range dnGroups {
		if strings.Contains(group, "LAR Redbook") {
			fmt.Println("controllers/adauth. Contains Redbook Auth")
			login(g, data.Username)
			approved = true
			break
		}
	}
	//If not in group send back unauthorised
	if !approved {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "Not in Approved Uploader Group"})
	}

	fmt.Println("controllers/adauth. before returning token")

}

// login is a handler that parses a form and checks for specific data
func login(c *gin.Context, username string) {
	session := sessions.Default(c)
	fmt.Println("controllers/adauth.go good auth, login called")

	// Save the username in the session
	session.Set(username, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		fmt.Println("controllers/adauth.go Failed to Save session")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	fmt.Println("controllers/adauth.go good auth, succesful ")

	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

func logout(c *gin.Context, username string) {
	session := sessions.Default(c)
	user := session.Get(username)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(username)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func me(c *gin.Context, username string) {
	session := sessions.Default(c)
	user := session.Get(username)
	fmt.Println("controllers/me session grabbed")
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}

/*
//ObjectGroups returns which of the given groups (referenced by DN) the object with the given attribute value is in,
//if any, or an error if one occurred.
func (c *Conn) ObjectGroups(attr, value string, groups []string) ([]string, error) {
	entry, err := c.GetAttributes(attr, value, []string{"memberOf"})
	if err != nil {
		return nil, err
	}
	var objectGroups []string

	for _, objectGroup := range entry.GetAttributeValues("memberOf") {
		for _, parentGroup := range groups {
			if objectGroup == parentGroup {
				objectGroups = append(objectGroups, parentGroup)
				continue
			}
		}
	}

	return objectGroups, nil
}

/*

//Connect returns an open connection to an Active Directory server or an error if one occurred.
func (c *Config) Connect() (*Conn, error) {
	switch c.Security {
	case SecurityNone:
		conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", c.Server, c.Port))
		if err != nil {
			return nil, fmt.Errorf("Connection error: %v", err)
		}
		return &Conn{Conn: conn, Config: c}, nil
	case SecurityTLS:
		conn, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", c.Server, c.Port), &tls.Config{ServerName: c.Server})
		if err != nil {
			return nil, fmt.Errorf("Connection error: %v", err)
		}
		return &Conn{Conn: conn, Config: c}, nil
	case SecurityStartTLS:
		conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", c.Server, c.Port))
		if err != nil {
			return nil, fmt.Errorf("Connection error: %v", err)
		}
		err = conn.StartTLS(&tls.Config{ServerName: c.Server})
		if err != nil {
			return nil, fmt.Errorf("Connection error: %v", err)
		}
		return &Conn{Conn: conn, Config: c}, nil
	default:
		return nil, errors.New("Configuration error: invalid SecurityType")
	}
}

//Bind authenticates the connection with the given userPrincipalName and password
//and returns the result or an error if one occurred.
func (c *Conn) Bind(upn, password string) (bool, error) {
	if password == "" {
		return false, nil
	}

	err := c.Conn.Bind(upn, password)
	if err != nil {
		if e, ok := err.(*ldap.Error); ok {
			if e.ResultCode == ldap.LDAPResultInvalidCredentials {
				return false, nil
			}
		}
		return false, fmt.Errorf("Bind error (%s): %v", upn, err)
	}

	return true, nil
}
*/
