package controllers

import (
	"net/http"
	"server/models"
	"server/types"
	"time"

	"github.com/gin-gonic/gin"
)

func createOrUpdateNews(g *gin.Context, m *models.Context) {
	var data types.News

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Check if all required parameters exist.
	if data.Author == "" || data.Date == "" || data.Text == "" || data.Title == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss one or multiple parameters"})
		return
	}

	// Check ID is not defined.
	if data.ID != 0 && g.Param("name") == "createNews" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "ID should not be defined"})
		return
	}

	// Check Date.
	t, err := time.Parse(time.RFC3339, data.Date)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Format StartDate.
	data.Date = t.Format(time.RFC3339)

	// Update News record; or create a new News record if ID is not defined.
	if err := m.CreateOrUpdateNews(&data); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Serve the result.
	g.JSON(http.StatusOK, data)
}
