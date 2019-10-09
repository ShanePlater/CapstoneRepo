package controllers

import (
	"net/http"
	"server/models"
	"server/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func getProjectsSiteInspectionsByProjectNumber(g *gin.Context, m *models.Context) {
	var data types.GetByIDJSON

	//B14453 Is a good testing project for site inspections
	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Return nil if ID is whitespace only or empty.
	if strings.Replace(data.ID, " ", "", -1) == "" {
		g.JSON(http.StatusBadRequest, nil)
		return
	}

	g.JSON(http.StatusOK, m.GetProjectsSiteInspectionsByProjectNumber(&data))
}