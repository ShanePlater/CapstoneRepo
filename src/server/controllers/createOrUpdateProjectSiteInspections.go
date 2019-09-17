package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func createOrUpdateProjectSiteInspections(g *gin.Context, m *models.Context) {
	fmt.Println("Attempting createOrUpdateProjectSiteInspections.go in controllers")

	var data types.ProjectsSiteInspections

	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	// client name, project number, pro name, address, suburb, location, type, status, start date, client rep name, telephone, mobile, email, division
	if data.InspectionID == "" || data.ProjectNumber == "" || data.InspectedBy == "" || data.InspectionDateTime == "" || data.InspectionDetails == "" {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss one or multiple parameters"})
		fmt.Println("controllers/createOrUpdateProjectSiteInspection.go missing a parameter error")
		//fmt.Println("\n name" + data.Name + "\n client" + data.ClientName + "\n address" + data.Address + "\n suburb" + data.Suburb + "\n location" + data.Location + "\n type" + data.Type + "\n status" + data.Status + "\n startdate" + data.StartDate + "\n enddate" + data.EndDate + "\n CRName" + data.CRName + "\n CRPhone" + data.CRPhone + "\n CRmobile" + data.CRMobile + "\n CREemail" + data.CREmail + "\n division" + data.Division)
		return
	}

	// Update project record; or create a new project record if ID is not defined.
	if err := m.CreateOrUpdateSiteInspections(&data); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		fmt.Println("controllers/createOrUpdateProjectSiteInspection.go  there was an error updateing or creating the ProjectSiteInspection")
		return
	}

	// Serve the result.
	g.JSON(http.StatusOK, data)
}