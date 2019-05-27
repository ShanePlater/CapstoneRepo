package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"
	"time"

	"github.com/gin-gonic/gin"
)

func createOrUpdateProjects(g *gin.Context, m *models.Context) {
	fmt.Println("Attempting createOrUpdateProjects.go in controllers")

	var data types.Project2

	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	// client name, project number, pro name, address, suburb, location, type, status, start date, client rep name, telephone, mobile, email, division
	if data.Name == "" || data.ClientName == "" || data.Address == "" || data.Suburb == "" || data.Location == "" || data.Type == "" || data.Status == "" || data.StartDate == "" || data.EndDate == "" || data.CRName == "" || data.CRPhone == "" || data.CRMobile == "" || data.CREmail == "" || data.Division == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss one or multiple parameters"})
		fmt.Println("controllers/createOrUpdateProjects.go shit broke fam, missing a parameter error")
		return
	}

	// Check Start Date Format.
	t, err := time.Parse(time.RFC3339, data.StartDate)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		fmt.Println("controllers/createOrUpdateProjects.go Error formatting the START date and time mate")
		return
	}
	// Check End Date Format.
	t, err = time.Parse(time.RFC3339, data.EndDate)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		fmt.Println("controllers/createOrUpdateProjects.go Error formatting the END date and time mate")
		return
	}

	// Format StartDate. dont think this is required, if it gets to this point then the dates should already be correct
	data.StartDate = t.Format(time.RFC3339)

	// Update project record; or create a new project record if ID is not defined.
	if err := m.CreateOrUpdateProjects(&data); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		fmt.Println("controllers/createOrUpdateProjects.go AHH fuck shit, there was an error updateing or creating the project")
		return
	}

	// Serve the result.
	g.JSON(http.StatusOK, data)
}
