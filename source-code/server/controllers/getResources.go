package controllers

import (
	"errors"
	"net/http"
	"server/models"
	"server/types"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func getResources(g *gin.Context, m *models.Context) {
	var data types.GetResourcesJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Situation that CategoryID is not set.
	if data.CategoryID == 0 {
		g.JSON(http.StatusBadRequest, gin.H{"Error": "Miss CategoryID"})
		return
	}

	// Situation if CategoryID or file path is not defined.
	if _, ok := m.GetCompanyDocumentResourceCategoriesTable().Load(utils.Itoa(data.CategoryID)); !ok {
		g.JSON(http.StatusInternalServerError, gin.H{"Error": errors.New("file path not found")})
		return
	}

	// Serve the result.
	// -1 is meant to ignore number of document records (get as more as possible).
	g.JSON(http.StatusOK, m.GetRecentCompanyDocumentResources(data.CategoryID, -1))
}
