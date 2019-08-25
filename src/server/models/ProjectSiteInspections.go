package models

import (
	"fmt"
	"server/orm"
	"server/types"
)

//CreateOrUpdateSiteInspections create client class and upload to database
func (c *Context) CreateOrUpdateSiteInspections(data *types.ProjectSiteInspections) error {
	fmt.Println("models/ CreatingOrUpdateSiteInspections")

	//Get division code, used Brisbane as the default office as we dont have proper AD sync to take it from user yet
	//data.Division = utils.Itoa(c.locateDivisionCode("BNE", data.Division))

	//current hard coded values: clientID, Client Office Code
	//Project Director & Manager --> The DB has 0 for many of these, after that it is the domain login, implemented in a later sprint.
	r := orm.ProjectSiteInspections{
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
		c.projectSiteInspections.Store(r.InspectionID, r)
	}

	// Update data.
	*data = types.ProjectSiteInspections{
		InspectionID:       r.InspectionID,
		ProjectNumber:      r.ProjectNumber,
		InspectedBy:        r.InspectedBy,
		InspectionDateTime: r.InspectionDateTime,
		InspectionDetails:  r.InspectionDetails,
	}

	return nil
}
