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
		//fmt.Println("\n name" + data.Name + "\n client" + data.ClientName + "\n address" + data.Address + "\n suburb" + data.Suburb + "\n location" + data.Location + "\n type" + data.Type + "\n status" + data.Status + "\n startdate" + data.StartDate + "\n enddate" + data.EndDate + "\n CRName" + data.CRName + "\n CRPhone" + data.CRPhone + "\n CRmobile" + data.CRMobile + "\n CREemail" + data.CREmail + "\n division" + data.Division)
		return
	}

	/*
		// Check Start Date Format.
		t, err := time.Parse(time.RFC3339, data.StartDate)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			fmt.Println("controllers/createOrUpdateProjects.go Error formatting the START date and time")
			return
		}
		// Check End Date Format.
		t, err = time.Parse(time.RFC3339, data.EndDate)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			fmt.Println("controllers/createOrUpdateProjects.go Error formatting the END date and time")
			return
		}


		// Format StartDate. dont think this is required, if it gets to this point then the dates should already be correct
		data.StartDate = t.Format(time.RFC3339)

		data.EndDate = t.Format(time.RFC3339)

	*/
	// Update project record; or create a new project record if ID is not defined.
	if err := m.CreateOrUpdateClients(&data); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		fmt.Println("controllers/createOrUpdateClients.go  there was an error updateing or creating the client")
		return
	}

	// Serve the result.
	g.JSON(http.StatusOK, data)
}