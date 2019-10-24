package controllers

import (
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func createSuggestion(g *gin.Context, m *models.Context) {
	var data types.CreateSuggestionJSON
	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Check if all required parameters exist.
	if data.Author == "" || data.Text == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss one or multiple parameters"})
		return
	}

	// Send email.
	if err := m.SendEmail(&data, types.SuggestionSubject, types.SuggestionTmpl); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err})
		return
	}

	// Serve the result.
	g.JSON(http.StatusOK, nil)
}
