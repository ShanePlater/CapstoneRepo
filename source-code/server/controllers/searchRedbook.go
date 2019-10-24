package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func searchRedbook(g *gin.Context, m *models.Context) {
	fmt.Println("Attempting Redbook Search")
	var data types.SearchRedbookJSON

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
	g.JSON(http.StatusOK, m.SearchRedbook(&data))
}
