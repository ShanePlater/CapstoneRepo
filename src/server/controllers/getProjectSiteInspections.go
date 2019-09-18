package controllers

import (
	"fmt"
	"net/http"
	"server/models"
	"server/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func getProjectSiteInspections(g *gin.Context, m *models.Context) {
	var data types.GetByIDJSON
	// Valid project w/ site inspection: B14533 Project Number for debugging
	fmt.Println("controllers/getProjectSiteInspections.go  Linked")
	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Return nil if ID is whitespace only or empty.
	if strings.Replace(data.ID, " ", "", -1) == "" {
		fmt.Println("controllers/getProjectSiteInspections.go Return Nil")
		g.JSON(http.StatusBadRequest, nil)
		return
	}
	fmt.Println("controllers/getProjectSiteInspections.go Past Bad Request")

	if res, ok := m.GetProjectSiteInspections(&data); ok {
		// Serve the result.
		fmt.Println("controllers/getProjectSiteInspections.go  Result Served")
		g.JSON(http.StatusOK, res)
		return
	} else {
		fmt.Println(ok)
	}
	g.JSON(http.StatusBadRequest, nil)
}
