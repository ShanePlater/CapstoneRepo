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
	fmt.Println("controllers/getProject.go  Linked")
	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	fmt.Println("controllers/getProject.go  Unmarshelled")

	// Return nil if ID is whitespace only or empty.
	if strings.Replace(data.ID, " ", "", -1) == "" {
		fmt.Println("controllers/getProject.go  whitespace ID")
		g.JSON(http.StatusBadRequest, nil)
		return
	}

	fmt.Println("controllers/getProject.go  past bad request")

	if res, ok := m.GetProject(&data); ok {
		// Serve the result.
		fmt.Println("controllers/getProject.go  Result served")
		g.JSON(http.StatusOK, res)
		return
	}
	g.JSON(http.StatusBadRequest, nil)
}
