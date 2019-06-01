package models

import (
	"server/config"
	"sync"

	"github.com/jinzhu/gorm"
)

// Context is the data structure utilized for cache and database operation.
// It allows models to be the abstract layer upon orm.
type Context struct {
	// Configuration.
	config *config.Context

	// Database.
	db *gorm.DB

	// News table.
	news *sync.Map

	// CompanyDocumentResources table.
	companyDocumentResources *sync.Map

	// CompanyDocumentResourceCategories table.
	companyDocumentResourceCategories *sync.Map

	// HelpDeskTicketAttachedFiles table.
	helpDeskTicketAttachedFiles *sync.Map

	// HelpDeskTickets table.
	helpDeskTickets *sync.Map

	// ProjectLocations table.
	projectLocations *sync.Map

	// ProjectTypes table.
	projectTypes *sync.Map

	// ProjectStatuss table.
	projectStatuss *sync.Map

	// Divisions table.
	divisions *sync.Map

	// Offices table.
	offices *sync.Map

	// Projects table.
	projects *sync.Map

	// Clients table.
	clients *sync.Map

	// Users table.
	users *sync.Map

	//clientlocationcode table
	clientlocations *sync.Map

	//clienttypecode table
	clienttypes *sync.Map
}

// copy is utilized to pass data from interface to another interface.
type copy struct {
	context *Context
	intf    []interface{}
	// Reserved positions for holding temporary data.
	reserveString []string
	reserveInt    []int
}
