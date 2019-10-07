package models

import (
	"fmt"
	"server/config"
	"server/orm"
)

// New returns an initialized model context with configurations read.
func New(config *config.Context) *Context {
	c := &Context{
		db:     orm.OpenDB(config.Get("mssql")),
		config: config,
	}

	// Check if cache is enabled.
	if !c.config.IsCache() {
		return c
	}

	// Store data in memory.
	c.news = orm.LoadNews(c.db)
	c.companyDocumentResources = orm.LoadCompanyDocumentResources(c.db)
	c.companyDocumentResourceCategories = orm.LoadCompanyDocumentResourceCategories(c.db)
	c.helpDeskTicketAttachedFiles = orm.LoadHelpDeskTicketAttachedFiles(c.db)
	c.helpDeskTickets = orm.LoadHelpDeskTickets(c.db)
	c.projectLocations = orm.LoadProjectLocations(c.db)
	c.projectTypes = orm.LoadProjectTypes(c.db)
	c.projectStatuss = orm.LoadProjectStatuss(c.db)
	c.divisions = orm.LoadDivisions(c.db)
	c.offices = orm.LoadOffices(c.db)
	c.projects = orm.LoadProjects(c.db)
	c.clients = orm.LoadClients(c.db)
	c.users = orm.LoadUsers(c.db)
	c.clientTypes = orm.LoadClientTypeCode(c.db)
	c.clientLocations = orm.LoadClientLocationCode(c.db)
	c.projectsSiteInspections = orm.LoadSiteInspections(c.db)
	fmt.Println("models/models.go  FIRST DATA STORAGE")

	return c
}
