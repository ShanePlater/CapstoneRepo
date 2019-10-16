package models

import (
	"fmt"
	"reflect"
	"server/orm"
	"server/types"
	"server/utils"
	"sort"
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
	if a.reserveString[0] != "" && !CaseInsensitiveContains(p.ClientName, a.reserveString[0]) && !CaseInsensitiveContains(p.Company, a.reserveString[0]) {
		// old strings.Contains(p.ClientName, a.reserveString[0]) && !strings.Contains(p.Company, a.reserveString[0])
		return true
	}

	// Check if address matches.
	if a.reserveString[1] != "" && !CaseInsensitiveContains(p.Address, a.reserveString[1]) {
		// old !strings.Contains(p.Address, a.reserveString[1])
		return true
	}

	// Check if suburb matches.
	if a.reserveString[2] != "" && !CaseInsensitiveContains(reflect.ValueOf(value).FieldByName("ProjectSuburb").String(), a.reserveString[2]) {
		//  strings.Contains(reflect.ValueOf(value).FieldByName("ProjectSuburb").String(), a.reserveString[2])
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
		if CaseInsensitiveContains(reflect.ValueOf(value).Field(i).String(), a.reserveString[0]) {
			//if strings.Contains(reflect.ValueOf(value).Field(i).String(), a.reserveString[0]) {
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

//CreateOrUpdateProjects comment
func (c *Context) CreateOrUpdateProjects(data *types.Project2) error {
	fmt.Println("models/ CreatingOrUpdateProjects")

	//Get division code, used Brisbane as the default office as we dont have proper AD sync to take it from user yet
	data.Division = utils.Itoa(c.locateDivisionCode("BNE", data.Division))

	//current hard coded values: clientID: I will fix this soon - Matt
	//Project Director & Manager --> The DB has 0 for many of these, after that it is the domain login, implemented in a later sprint.
	r := orm.Projects{
		ProjectNumber:         data.ID,
		ProjectName:           data.Name,
		ClientID:              "3065",
		Division:              data.Division,
		ClientRepName:         data.CRName,
		ClientRepTelephone:    data.CRPhone,
		ClientRepMobile:       data.CRMobile,
		ClientRepEmailAddress: data.CREmail,
		ProjectDirector:       "0",
		ProjectManager:        "0",
		ProjectStatusCode:     data.Status,
		ProjectStartDate:      data.StartDate,
		ProjectEndDate:        data.EndDate,
		ProjectTypeCode:       data.Type,
		ProjectAddress:        data.Address,
		ProjectSuburb:         data.Suburb,
		ProjectLocationCode:   data.Location,
		ProjectDescription:    data.Description,
		ProjectValue:          data.Value,
		ArchiveLocation:       data.ArchiveLocation,
	}

	if err := orm.CreateOrUpdateProjects(&r, c.db); err != nil {
		fmt.Println("Models/ Error in Projects.go")
		return err
	}

	// Check if cache is enabled.
	if c.config.IsCache() {
		c.projects.Store(r.ProjectNumber, r)
	}

	// Update data. try removing this and see if it e
	*data = types.Project2{
		ID:              r.ProjectNumber,
		Name:            r.ProjectName,
		Address:         r.ProjectAddress,
		Suburb:          r.ProjectSuburb,
		Location:        r.ProjectLocationCode,
		Type:            r.ProjectTypeCode,
		Status:          r.ProjectStatusCode,
		StartDate:       r.ProjectStartDate,
		EndDate:         r.ProjectEndDate,
		CRName:          r.ClientRepName,
		CRPhone:         r.ClientRepTelephone,
		CRMobile:        r.ClientRepMobile,
		CREmail:         r.ClientRepEmailAddress,
		Division:        r.Division,
		Director:        r.ProjectDirector,
		Manager:         r.ProjectManager,
		Value:           r.ProjectValue,
		Description:     r.ProjectDescription,
		ArchiveLocation: r.ArchiveLocation,
	}

	return nil
}
