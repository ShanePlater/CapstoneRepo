package controllers

import (
	"net/http"
	"server/models"
	"server/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func getSuggestionItemURL(g *gin.Context, m *models.Context) {
	var data types.SearchJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Return nil if keyword is whitespace only or empty.
	if strings.Replace(data.Keyword, " ", "", -1) == "" {
		g.JSON(http.StatusBadRequest, nil)
		return
	}

	// Serve the result.
	g.JSON(http.StatusOK, gin.H{"URL": m.GetSuggestionItemURL(&data)})
}
