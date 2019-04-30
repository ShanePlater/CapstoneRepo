package controllers

import (
	"net/http"
	"server/models"

	"github.com/gin-gonic/gin"
)

func getHelpDeskTickets(g *gin.Context, m *models.Context) {
	// Serve the result.
	g.JSON(http.StatusOK, m.GetHelpDeskTickets())
}
