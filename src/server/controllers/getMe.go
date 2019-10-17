package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"

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

	status, err := auth.Authenticate(config, data.Username, data.Password)
	//If authentication was successful then save the session

	if status == true {
		fmt.Println("controllers/adauth.go good auth, prelogin")
		login(g, data.Username)
	}
	if err != nil {
		fmt.Println("controllers/adauth.go error when authenticating")
		fmt.Println(err)
		return
	}
	if !status {
		//handle failed authentication
		g.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	fmt.Println("controllers/adauth. before returning token")
	//token := md5.New()
	//token.Write([]byte(data.Username + strconv.FormatInt(time.Now().Unix(), 10) + data.Password))
	// Serve the result.
	//g.JSON(http.StatusOK, gin.H{"Token": hex.EncodeToString(token.Sum(nil))})
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

// AuthRequired is a simple middleware to check the session,
func AuthRequired(c *gin.Context) {
	var data types.GetMeJSON

	// Unmarshal application/json and bind to struct.
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	username := data.Username
	session := sessions.Default(c)
	user := session.Get(username)
	if user == nil {
		// Abort the request with the appropriate error code
		fmt.Println("controllers/AuthRequired.go user == nil")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	username = "Success"
	// Continue down the chain to handler etc
	fmt.Println("controllers/AuthRequired.go past user verification")
	c.JSON(http.StatusOK, username)
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
