package controllers

import (
	"net/http"
	"server/models"
	"server/types"
	"time"

	"github.com/gin-gonic/gin"
)

func createOrUpdateProjects(g *gin.Context, m *models.Context) {
	var data types.Project2

	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	// client name, project number, pro name, address, suburb, location, type, status, start date, client rep name, telephone, mobile, email, division
	if data.Name == "" || data.ClientName == "" || data.Address == "" || data.Suburb == "" || data.Location == "" || data.Type == "" || data.Status == "" || data.StartDate == "" || EndDate == "" || data.CRName == "" || data.CRPhone == "" || data.CRMobile == "" || data.CREmail == "" || data.Division == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss one or multiple parameters"})
		return
	}

	// Check Start Date Format.
	t, err := time.Parse(time.RFC3339, data.StartDate)	
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	// Check End Date Format.
	t, err := time.Parse(time.RFC3339, data.EndDate)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}	

	// Format StartDate. dont think this is required, if it gets to this point then the dates should already be correct
	//data.StartDate = t.Format(time.RFC3339)

	// Update project record; or create a new project record if ID is not defined.
	if err := m.CreateOrUpdateProjects(&data); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Serve the result.
	g.JSON(http.StatusOK, data)
}
