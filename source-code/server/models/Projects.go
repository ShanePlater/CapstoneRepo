package models

import (
	"reflect"
	"server/orm"
	"server/types"
	"sort"
	"strings"
)

// SearchRedbook return projects which achieve the given requirements.
func (c *Context) SearchRedbook(data *types.SearchRedbookJSON) interface{} {
	// Set given time period.
	res := &copy{
		context: c,
		reserveString: []string{
			trimUnexpectedCharactors(data.ClientName),
			trimUnexpectedCharactors(data.Address),
			trimUnexpectedCharactors(data.Suburb),
			data.Location,
			data.Type,
			data.Status,
			data.Division,
			data.Office,
			data.StartDate,
			data.EndDate}}

	// Obtain records in time period.
	c.GetProjectsTable().Range(res.rangeRedbook)

	// Sort records in order latest to oldest.
	sort.Sort(byID(res.intf))

	return res.intf
}

// rangeRedbook pass records which achieve requirements.
func (a *copy) rangeRedbook(key, value interface{}) bool {
	p := types.Project{
		ID:         reflect.ValueOf(value).FieldByName("ProjectNumber").String(),
		Name:       reflect.ValueOf(value).FieldByName("ProjectName").String(),
		ClientName: reflect.ValueOf(value).FieldByName("ClientRepName").String(),
		Address:    reflect.ValueOf(value).FieldByName("ProjectAddress").String(),
	}

	// Check if start date is in time period.
	if reflect.ValueOf(value).FieldByName("ProjectStartDate").String() != "" && !inTimePeriod(a.reserveString[8], a.reserveString[9], reflect.ValueOf(value).FieldByName("ProjectStartDate").String()) {
		return true
	}

	// Check if end date is in time period.
	if reflect.ValueOf(value).FieldByName("ProjectEndDate").String() != "" && !inTimePeriod(a.reserveString[8], a.reserveString[9], reflect.ValueOf(value).FieldByName("ProjectEndDate").String()) {
		return true
	}

	// Load client entity based on ClientID.
	client, ClientExist := a.context.GetClientsTable().Load(reflect.ValueOf(value).FieldByName("ClientID").String())

	if !ClientExist {
		// Client is not found.
		// ClientID in Projects table point to the non-existent client entity.
		p.Company = ""
	} else {
		p.Company = reflect.ValueOf(client).FieldByName("ClientName").String()
	}

	// Check if client name or company name matches.
	if a.reserveString[0] != "" && !strings.Contains(p.ClientName, a.reserveString[0]) && !strings.Contains(p.Company, a.reserveString[0]) {
		return true
	}

	// Check if address matches.
	if a.reserveString[1] != "" && !strings.Contains(p.Address, a.reserveString[1]) {
		return true
	}

	// Check if suburb matches.
	if a.reserveString[2] != "" && !strings.Contains(reflect.ValueOf(value).FieldByName("ProjectSuburb").String(), a.reserveString[2]) {
		return true
	}

	// Check if location matches.
	if a.reserveString[3] != "" && a.reserveString[3] != reflect.ValueOf(value).FieldByName("ProjectLocationCode").String() {
		return true
	}

	// Check if type matches.
	if a.reserveString[4] != "" && a.reserveString[4] != reflect.ValueOf(value).FieldByName("ProjectTypeCode").String() {
		return true
	}

	// Check if status matches.
	if a.reserveString[5] != "" && a.reserveString[5] != reflect.ValueOf(value).FieldByName("ProjectStatusCode").String() {
		return true
	}

	// Load division entity based on Division.
	division, divisionExist := a.context.GetDivisionsTable().Load(reflect.ValueOf(value).FieldByName("Division").String())

	if !divisionExist {
		// Division is not found.
		// Divison in Projects table point to the non-existent division entity.
		p.Division = ""

		// There is no office information because division is not found.
		if a.reserveString[7] != "" {
			// Skip if office is set for search.
			return true
		}
	} else {
		p.Division = reflect.ValueOf(division).FieldByName("DivisionName").String()

		// Check if office matches.
		if a.reserveString[7] != "" && a.reserveString[7] != reflect.ValueOf(division).FieldByName("OfficeCode").String() {
			return true
		}
	}

	if a.reserveString[6] != "" && a.reserveString[6] != p.Division {
		return true
	}

	// Pass value to interface in copy.
	a.intf = append(a.intf, p)

	// Return true for keep looping through map. Otherwise, loop exit.
	return true
}

// SearchProjects return projects which achieve the given requirements.
func (c *Context) SearchProjects(keyword string) interface{} {
	// Set keyword for search.
	res := &copy{context: c,
		reserveString: []string{keyword}}

	// Obtain records which contains the keyword.
	c.GetProjectsTable().Range(res.rangeProjectsSearch)

	// Sort records in order latest to oldest.
	sort.Sort(byID(res.intf))

	return res.intf
}

// rangeProjectsSearch pass records which achieve requirements.
func (a *copy) rangeProjectsSearch(key, value interface{}) bool {
	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		if strings.Contains(reflect.ValueOf(value).Field(i).String(), a.reserveString[0]) {
			p := types.Project{
				ID:         reflect.ValueOf(value).FieldByName("ProjectNumber").String(),
				Name:       reflect.ValueOf(value).FieldByName("ProjectName").String(),
				ClientName: reflect.ValueOf(value).FieldByName("ClientRepName").String(),
				Address:    reflect.ValueOf(value).FieldByName("ProjectAddress").String(),
			}

			// Load client entity based on ClientID.
			if c, ok := a.context.GetClientsTable().Load(reflect.ValueOf(value).FieldByName("ClientID").String()); ok {
				p.Company = reflect.ValueOf(c).FieldByName("ClientName").String()
			}

			// Load division entity based on Division.
			if d, ok := a.context.GetDivisionsTable().Load(reflect.ValueOf(value).FieldByName("Division").String()); ok {
				p.Division = reflect.ValueOf(d).FieldByName("DivisionName").String()
			}

			// Pass value to interface in copy.
			a.intf = append(a.intf, p)

			return true
		}
	}
	return true
}
func (c *Context) CreateOrUpdateProjects(data *types.Project) error {
	r := orm.Projects{
		// ArticleID:       utils.Itoa(data.ID),
		// PostedBy:        data.Author,
		// DateTimePosted:  data.Date,
		// ArticleHeadline: data.Title,
		// ArticleText:     data.Text,
	}

	if err := orm.CreateOrUpdateProjects(&r, c.db); err != nil {
		return err
	}

	// Check if cache is enabled.
	if c.config.IsCache() {
		c.projects.Store(r.ProjectNumber, r)
	}

	// Update data.
	*data = types.Project{
		// ID:     utils.Atoi(r.ArticleID),
		// Author: r.PostedBy,
		// Date:   r.DateTimePosted,
		// Title:  r.ArticleHeadline,
		// Text:   r.ArticleText,
	}

	return nil
}
