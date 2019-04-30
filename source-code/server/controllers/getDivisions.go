package controllers

import (
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func getDivisions(g *gin.Context, m *models.Context) {
	var data types.GetDivisionsJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Return nil if ID is whitespace only or empty.
	if data.Office == "" || data.Division == "" {
		g.JSON(http.StatusBadRequest, nil)
		return
	}

	g.JSON(http.StatusOK, m.GetUsersByDivisionCode(&data))
}
