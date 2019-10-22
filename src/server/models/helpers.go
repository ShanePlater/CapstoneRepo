package models

import (
	"fmt"
	"server/config"
	"server/orm"
	"server/types"
	"sync"
)

// GetConf exports configuration.
func (c *Context) GetConf() *config.Context {
	return c.config
}

// GetNewsTable exports News table.
func (c *Context) GetNewsTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.news
	}

	// Load data from database then return it.
	return orm.LoadNews(c.db)
}

//GetClientTypesTable gets Client Types.
func (c *Context) GetClientTypesTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.news
	}

	// Load data from database then return it.
	return orm.LoadClientTypeCode(c.db)
}

// GetClientLocationsTable exports ClientLocations Table.
func (c *Context) GetClientLocationsTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.news
	}

	// Load data from database then return it.
	return orm.LoadClientLocationCode(c.db)
}

// GetCompanyDocumentResourcesTable exports CompanyDocumentResources table.
func (c *Context) GetCompanyDocumentResourcesTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.companyDocumentResources
	}

	// Load data from database then return it.
	return orm.LoadCompanyDocumentResources(c.db)
}

// GetCompanyDocumentResourceCategoriesTable exports CompanyDocumentResourceCategories table.
func (c *Context) GetCompanyDocumentResourceCategoriesTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.companyDocumentResourceCategories
	}

	// Load data from database then return it.
	return orm.LoadCompanyDocumentResourceCategories(c.db)
}

// GetHelpDeskTicketAttachedFilesTable exports HelpDeskTicketAttachedFiles table.
func (c *Context) GetHelpDeskTicketAttachedFilesTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.helpDeskTicketAttachedFiles
	}

	// Load data from database then return it.
	return orm.LoadHelpDeskTicketAttachedFiles(c.db)
}

// GetHelpDeskTicketsTable exports News table.
func (c *Context) GetHelpDeskTicketsTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.helpDeskTickets
	}

	// Load data from database then return it.
	return orm.LoadHelpDeskTickets(c.db)
}

// GetProjectLocationsTable exports ProjectLocations table.
func (c *Context) GetProjectLocationsTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.projectLocations
	}

	// Load data from database then return it.
	return orm.LoadProjectLocations(c.db)
}

// GetProjectTypesTable exports ProjectTypes table.
func (c *Context) GetProjectTypesTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.projectTypes
	}

	// Load data from database then return it.
	return orm.LoadProjectTypes(c.db)
}

// GetProjectStatussTable exports ProjectStatuss table.
func (c *Context) GetProjectStatussTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.projectStatuss
	}

	// Load data from database then return it.
	return orm.LoadProjectStatuss(c.db)
}

// GetDivisionsTable exports Divisions table.
func (c *Context) GetDivisionsTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.divisions
	}

	// Load data from database then return it.
	return orm.LoadDivisions(c.db)
}

// GetOfficesTable exports Offices table.
func (c *Context) GetOfficesTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.offices
	}

	// Load data from database then return it.
	return orm.LoadOffices(c.db)
}

// GetProjectsTable exports Projects table.
func (c *Context) GetProjectsTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.projects
	}

	// Load data from database then return it.
	return orm.LoadProjects(c.db)
}

// GetClientsTable exports Clients table.
func (c *Context) GetClientsTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.clients
	}

	// Load data from database then return it.
	return orm.LoadClients(c.db)
}

// GetUsersTable exports Users table.
func (c *Context) GetUsersTable() *sync.Map {
	if c.config.IsCache() {
		// Return from cache.
		return c.users
	}

	// Load data from database then return it.
	return orm.LoadUsers(c.db)
}

// GetProjectSiteInspectionsTable exports ProjectSiteInspections table.
func (c *Context) GetProjectSiteInspectionsTable() *sync.Map {
	if c.config.IsCache() {
		fmt.Println("controllers/getProjectSiteInspections.go  whygodwhy")
		return c.projectsSiteInspections
	}
	fmt.Println("controllers/getProjectSiteInspections.go  Cache enabled")
	// Load data from database then return it.
	return orm.LoadSiteInspections(c.db)
}

//ReloadCache reloads all tables into sync.map
func (c *Context) ReloadCache() *sync.Map {

	fmt.Println("controllers/getProjectSiteInspections.go  Reloading Tables")
	// Load data from database then return it.
	m := orm.LoadNews(c.db)

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

	fmt.Println("controllers/getProjectSiteInspections.go  Past Table Reloading")
	return m
}

//SwapProjLocCode returns the location code for the location name
func (c *Context) SwapProjLocCode(data *types.NameForCodes) *types.NameForCodes {
	var names *types.NameForCodes
	// Get each of the tables required
	// then get the code-name pair and return the NAMES in the data structure
	// this will then be pushed back up to controllers and returned to the front end
	/*
		projloccodes, err := c.GetProjectLocationsTable().Load(data.ProjectLocationCode)
		if err != true {
			fmt.Println(err)
		}
	*/

	/*
		The things we want to be returning
		names.ProjectLocationCode = corrosponding code name
		names.ProjectStatusCode = corro names
		names.ProjectTypeCode = corro name
	*/
	return names
}
