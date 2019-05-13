package controllers

import (
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func search(g *gin.Context, m *models.Context) {
	var data types.SearchJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Check if keyword is empty.
	if data.Keyword == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Cannot search without keyword."})
		return
	}

	// Serve the result.
	switch g.Param("name") {
	case "searchProjects":
		g.JSON(http.StatusOK, m.SearchProjects(data.Keyword))
	case "searchUsers":
		g.JSON(http.StatusOK, m.SearchUsers(data.Keyword))
	case "searchClients":
		g.JSON(http.StatusOK, m.SearchClients(data.Keyword))
	case "searchResources":
		g.JSON(http.StatusOK, m.SearchResources(data.Keyword))
	default:
	}
}
