package controllers

import (
	"net/http"
	"server/models"
	"server/types"

	"github.com/gin-gonic/gin"
)

func recent(g *gin.Context, m *models.Context) {
	var data types.RecentJSON

	// Unmarshal application/json and bind to struct.
	if err := g.BindJSON(&data); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// Situation that length is not set.
	if data.Length == 0 {
		data.Length = 3
	}

	// Serve the result.
	// -1 is meant to ignore categoryID (loop all records).
	g.JSON(http.StatusOK, m.GetRecentCompanyDocumentResources(-1, data.Length))
}
