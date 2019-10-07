package models

import (
	"fmt"
	"reflect"
	"server/orm"
	"server/types"
	"sort"
	"strings"
)

//CreateOrUpdateSiteInspections create client class and upload to database
func (c *Context) CreateOrUpdateSiteInspections(data *types.ProjectsSiteInspections) error {
	fmt.Println("models/ CreatingOrUpdateSiteInspections")

	//Get division code, used Brisbane as the default office as we dont have proper AD sync to take it from user yet
	//data.Division = utils.Itoa(c.locateDivisionCode("BNE", data.Division))

	//current hard coded values: clientID, Client Office Code
	//Project Director & Manager --> The DB has 0 for many of these, after that it is the domain login, implemented in a later sprint.
	r := orm.Projectssiteinspections{
		InspectionID:       "0",
		ProjectNumber:      data.ProjectNumber,
		InspectedBy:        data.InspectedBy,
		InspectionDateTime: data.InspectionDateTime,
		InspectionDetails:  data.InspectionDetails,
	}

	if err := orm.CreateOrUpdateProjectSiteInspections(&r, c.db); err != nil {
		fmt.Println("Models/ Error in SiteInspections.go")
		return err
	}

	// Check if cache is enabled.
	if c.config.IsCache() {
		c.projectsSiteInspections.Store(r.InspectionID, r)
	}

	// Update data.
	*data = types.ProjectsSiteInspections{
		InspectionID:       r.InspectionID,
		ProjectNumber:      r.ProjectNumber,
		InspectedBy:        r.InspectedBy,
		InspectionDateTime: r.InspectionDateTime,
		InspectionDetails:  r.InspectionDetails,
	}

	return nil
}

// GetProjectsSiteInspectionsByProjectNumber return inspections which achieve the given requirements
func (c *Context) GetProjectsSiteInspectionsByProjectNumber(data *types.GetByIDJSON) interface{} {
	// Set ID for search.
	res := &copy{context: c,
		reserveString: []string{data.ID}}

	// Obtain records which contains the ID.
	c.GetProjectSiteInspectionsTable().Range(res.rangeProjectsSiteInspectionsByProjectNumber)

	// Sort records in order latest to oldest.
	sort.Sort(byDate(res.intf))

	return res.intf
}

func (a *copy) rangeProjectsSiteInspectionsByProjectNumber(key, value interface{}) bool {
	if strings.Compare(reflect.ValueOf(value).FieldByName("ProjectNumber").String(), a.reserveString[0]) == 0 {
		p := types.ProjectsSiteInspections{
			InspectionID:       reflect.ValueOf(value).FieldByName("InspectionID").String(),
			InspectionDateTime: reflect.ValueOf(value).FieldByName("InspectionDateTime").String(),
			InspectionDetails:  reflect.ValueOf(value).FieldByName("InspectionDetails").String(),
		}
		if d, ok := a.context.GetUsersTable().Load(reflect.ValueOf(value).FieldByName("InspectedBy").String()); ok {
			p.InspectedBy = reflect.ValueOf(d).FieldByName("FirstName").String() + " " + reflect.ValueOf(d).FieldByName("LastName").String()
		} else {
			p.InspectedBy = reflect.ValueOf(value).FieldByName("InspectedBy").String()
		}
		// Pass value to interface in copy.
		a.intf = append(a.intf, p)
	}

	return true
}
