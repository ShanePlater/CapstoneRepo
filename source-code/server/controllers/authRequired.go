package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthRequired is a simple middleware to check the session,
func authRequired(c *gin.Context, m *models.Context) {
	var data types.GetMeJSON

	// Unmarshal application/json and bind to struct.
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	Username := data.Username
	session := sessions.Default(c)
	user := session.Get(Username)
	if user == nil {
		// Abort the request with the appropriate error code
		fmt.Println("controllers/AuthRequired.go user == nil")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	data.Username = "Success"

	// Continue down the chain to handler etc
	fmt.Println("controllers/AuthRequired.go past user verification")
	c.JSON(http.StatusOK, data)
}
