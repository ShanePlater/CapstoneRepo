package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func createOrUpdateClients(g *gin.Context, m *models.Context) {
	fmt.Println("Attempting createOrUpdateClients.go in controllers")

	var data types.Client2

	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	// client name, project number, pro name, address, suburb, location, type, status, start date, client rep name, telephone, mobile, email, division
	if data.ClientOfficeCode == "" || data.ClientName == "" || data.ClientABNNumber == "" || data.ClientACNNumber == "" || data.ClientTypeCode == "" || data.FirstName == "" || data.LastName == "" || data.ClientLocationCode == "" || data.StreetAddress == "" || data.StreetSuburb == "" || data.StreetPostcode == "" || data.PhoneNumber == "" || data.FaxNumber == "" || data.EMailAddress == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss one or multiple parameters"})
		fmt.Println("controllers/createOrUpdateClients.go missing a parameter error")
		fmt.Println("Client Office Code \n" + data.ClientOfficeCode + "Client Name \n" + data.ClientName)
		return
	}

	// Update project record; or create a new project record if ID is not defined.
	if err := m.CreateOrUpdateClients(&data); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		fmt.Println("controllers/createOrUpdateClients.go  there was an error updateing or creating the client")
		return
	}

	//Refresh Cache so Clients can be searched immediately after upload
	m.ReloadCache()
	// Serve the result.
	g.JSON(http.StatusOK, data)
}
