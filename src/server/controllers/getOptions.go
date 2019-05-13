package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func getOptions(g *gin.Context, m *models.Context) {
	// Serve the result.
	g.JSON(http.StatusOK, m.GetOptions(g.Param("name")))
}
