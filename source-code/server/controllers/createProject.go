package controllers

import (
	"net/http"
	"reflect"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func createProject(g *gin.Context, m *models.Context) {
	var data types.createProject
	// Unmarshal application/json and bind to struct.
	if err:= g.BindJSON(&data); err != nil{
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if data.Name == "" || data.ClientName == "" || data.Company == "" || data.Address == "" || data.Division == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "One or more fields have not been entered"})
		return
	}

	if err := m.CreateProject(&data); err != nil{
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	
	g.JSON(http.StatusOK, data)
}