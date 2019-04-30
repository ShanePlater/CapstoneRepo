package controllers

import (
	"net/http"
	"reflect"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func createHelpDeskTicket(g *gin.Context, m *models.Context) {
	var data types.HelpDeskTicket
	var files []types.File

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Check if all required parameters exist.
	if data.Situation == 0 || data.Username == "" || data.Computer == "" || data.Description == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss one or multiple parameters"})
		return
	}

	// Check ID and status is not defined.
	if data.TicketID != 0 || data.Status != 0 {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "ID and status should not be defined"})
		return
	}

	// Create help desk ticket and update data.
	if err := m.CreateHelpDeskTicket(&data); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// Update attached files information.
	for _, v := range data.Files {
		if err := m.UpdateAttachedFile(&v, data.TicketID); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		files = append(files, v)
	}

	// Put files slice into data.
	reflect.ValueOf(data).FieldByName("Files").Set(reflect.ValueOf(files))

	// Serve the result.
	g.JSON(http.StatusOK, data)
}
