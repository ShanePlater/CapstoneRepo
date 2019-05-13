package controllers

import (
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func getNews(g *gin.Context, m *models.Context) {
	var data types.GetNewsJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Parse time.
	if err := parseTime(&data.StartDate, &data.EndDate); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Serve the result.
	g.JSON(http.StatusOK, m.GetNews(&data))
}
