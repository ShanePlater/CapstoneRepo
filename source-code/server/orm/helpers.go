package orm

import (
	"sync"

	"github.com/jinzhu/gorm"
)

// LoadNews return all records in News table.
func LoadNews(db *gorm.DB) *sync.Map {
	var r []News

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.ArticleID, v)
	}

	return m
}

// LoadCompanyDocumentResources return all records in CompanyDocumentResources table.
func LoadCompanyDocumentResources(db *gorm.DB) *sync.Map {
	var r []Companydocumentresources

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.FileName, v)
	}

	return m
}

// LoadCompanyDocumentResourceCategories return all records in CompanyDocumentResourceCategories table.
func LoadCompanyDocumentResourceCategories(db *gorm.DB) *sync.Map {
	var r []Companydocumentresourcecategories

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.CategoryID, v)
	}

	return m
}

// LoadHelpDeskTicketAttachedFiles return all records in HelpDeskTicketAttachedFiles table.
func LoadHelpDeskTicketAttachedFiles(db *gorm.DB) *sync.Map {
	var r []Helpdeskticketattachedfiles

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.FileID, v)
	}

	return m
}

// LoadHelpDeskTickets return all records in HelpDeskTickets table.
func LoadHelpDeskTickets(db *gorm.DB) *sync.Map {
	var r []Helpdesktickets

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.TicketID, v)
	}

	return m
}

// LoadProjectLocations return all records in ProjectLocations table.
func LoadProjectLocations(db *gorm.DB) *sync.Map {
	var r []Projectlocations

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.ProjectLocationCode, v)
	}

	return m
}

// LoadProjectTypes return all records in ProjectTypes table.
func LoadProjectTypes(db *gorm.DB) *sync.Map {
	var r []Projecttypes

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.ProjectTypeCode, v)
	}

	return m
}

// LoadProjectStatuss return all records in ProjectStatuss table.
func LoadProjectStatuss(db *gorm.DB) *sync.Map {
	var r []Projectstatuss

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.ProjectStatusCode, v)
	}

	return m
}

// LoadDivisions return all records in Divisions table.
func LoadDivisions(db *gorm.DB) *sync.Map {
	var r []Divisions

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.DivisionCode, v)
	}

	return m
}

// LoadOffices return all records in Offices table.
func LoadOffices(db *gorm.DB) *sync.Map {
	var r []Offices

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.OfficeCode, v)
	}

	return m
}

// LoadProjects return all records in Projects table.
func LoadProjects(db *gorm.DB) *sync.Map {
	var r []Projects

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.ProjectNumber, v)
	}

	return m
}

// LoadClients return all records in Clients table.
func LoadClients(db *gorm.DB) *sync.Map {
	var r []Clients

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.ClientID, v)
	}

	return m
}

// LoadUsers return all records in Users table.
func LoadUsers(db *gorm.DB) *sync.Map {
	var r []Users

	// Get all records and check error.
	checkErr(db.Find(&r).Error)

	// Store records in sync.Map.
	m := &sync.Map{}
	for _, v := range r {
		m.Store(v.Username, v)
	}

	return m
}
