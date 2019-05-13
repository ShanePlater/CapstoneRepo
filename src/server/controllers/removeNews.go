package controllers

import (
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func removeNews(g *gin.Context, m *models.Context) {
	var data types.News

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Check if ID is missing
	if data.ID <= 0 {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Entity ID is required"})
		return
	}

	// Remove News record with given ID.
	if err := m.RemoveNews(&data.ID); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Serve the result.
	g.JSON(http.StatusOK, nil)
}
