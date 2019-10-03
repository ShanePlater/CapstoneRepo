package controllers

import (
	"net/http"
	"reflect"
	"server/models"
	"sync"

	"github.com/gin-gonic/gin"
)

// New returns an initialized controllers instance.
func New(m *models.Context) *Context {
	c := &Context{
		model:   m,
		handler: &sync.Map{},
	}

	// Store all handler functions to sync.Map handler.
	c.handler.Store("getNews", getNews)
	c.handler.Store("createNews", createOrUpdateNews)
	c.handler.Store("updateNews", createOrUpdateNews)
	c.handler.Store("createOrUpdateProject", createOrUpdateProjects)
	c.handler.Store("createOrUpdateClients", createOrUpdateClients)
	c.handler.Store("removeNews", removeNews)
	c.handler.Store("recent", recent)
	c.handler.Store("createSuggestion", createSuggestion)
	c.handler.Store("getResources", getResources)
	c.handler.Store("createHelpDeskTicket", createHelpDeskTicket)
	c.handler.Store("getHelpDeskTickets", getHelpDeskTickets)
	c.handler.Store("uploadFile", uploadFile)
	c.handler.Store("uploadResource", uploadResource)
	c.handler.Store("getOptionLocations", getOptions)
	c.handler.Store("getOptionTypes", getOptions)
	c.handler.Store("getOptionStatuss", getOptions)
	c.handler.Store("getOptionDivisions", getOptions)
	c.handler.Store("getOptionOffices", getOptions)
	c.handler.Store("getOptionCategories", getOptions)
	c.handler.Store("searchRedbook", searchRedbook)
	c.handler.Store("searchProjects", search)
	c.handler.Store("searchUsers", search)
	c.handler.Store("searchClients", search)
	c.handler.Store("searchResources", search)
	c.handler.Store("fetchSuggestions", fetchSuggestions)
	c.handler.Store("getClient", getClient)
	c.handler.Store("getProject", getProject)
	c.handler.Store("getSuggestionItemURL", getSuggestionItemURL)
	c.handler.Store("getDivisions", getDivisions)
	c.handler.Store("getMe", getMe)

	return c

}

// Serve determines which API is called and pass to the relevant API handler.
func Serve(c *Context, g *gin.Context) {
	// Check if request API exists.
	if f, ok := c.handler.Load(g.Param("name")); ok {
		// Call the function f with the input arguments in.
		reflect.ValueOf(f).Call([]reflect.Value{reflect.ValueOf(g), reflect.ValueOf(c.model)})
		return
	}
	// Otherwise, response as API not found.
	g.JSON(http.StatusNotFound, gin.H{
		"Error": "API not found",
	})
}
