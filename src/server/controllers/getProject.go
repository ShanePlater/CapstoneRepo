package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func getProject(g *gin.Context, m *models.Context) {
	var data types.GetByIDJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Return nil if ID is whitespace only or empty.
	if strings.Replace(data.ID, " ", "", -1) == "" {
		fmt.Println("controllers/getProject.go  whitespace ID")
		g.JSON(http.StatusBadRequest, nil)
		return
	}

	if res, ok := m.GetProject(&data); ok {
		// Serve the result.
		g.JSON(http.StatusOK, res)
		return
	}
	g.JSON(http.StatusBadRequest, nil)
}

func getProjectsByClientID(g *gin.Context, m *models.Context) {
	var data types.GetByIDJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Return nil if ID is whitespace only or empty.
	if strings.Replace(data.ID, " ", "", -1) == "" {
		fmt.Println("controllers/getProject.go  whitespace ID")
		g.JSON(http.StatusBadRequest, nil)
		return
	}
	res := m.GetProjectsByClientID(&data)
	g.JSON(http.StatusOK, res)
}
