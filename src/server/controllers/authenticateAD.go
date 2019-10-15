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

//ADAuth - Connects to the LAR domain when an authenticator.
func authenticateAD(g *gin.Context, m *models.Context) {

	var data types.Authenticator
	fmt.Println("controllers/adauth.go at controllers")
	if err := g.BindJSON(&data); err != nil {
		fmt.Println("controllers/adauth.go error when marshelling")
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	//Setup AD connection config
	config := &auth.Config{
		Server:   "LR-BNE-ERP-01",       /* Box that contains domain controller */
		Port:     389,                   /* Default Port for AD connection */
		BaseDN:   "DC=lar,DC=net,DC=au", /* Domain */
		Security: auth.SecurityNone,
	}

	username := data.Username
	password := data.Password

	status, err := auth.Authenticate(config, username, password)
	//If authentication was successful then save the session
	if status == true {
		fmt.Println("controllers/adauth.go good auth, prelogin")
		login(g, username, password)
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
}

// AuthRequired is a simple middleware to check the session,
func AuthRequired(c *gin.Context, username string) {
	session := sessions.Default(c)
	user := session.Get(username)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

// login is a handler that parses a form and checks for specific data
func login(c *gin.Context, username string, password string) {
	session := sessions.Default(c)
	fmt.Println("controllers/adauth.go good auth, login called")

	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Save the username in the session
	session.Set(username, username) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		fmt.Println("controllers/adauth.go Failed to Save session")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	fmt.Println("controllers/adauth.go good auth, succesful ")

	me(c, username)
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
