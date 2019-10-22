package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

//CreateOrUpdateProjectSiteInspections facilitating creation of new SiteInspections
func CreateOrUpdateProjectSiteInspections(p *Projectssiteinspections, db *gorm.DB) error {
	fmt.Println("orm/projectsiteinspections.go/ to create site inspection in orm")
	// Update record.
	/*
		if p.InspectionID != "0" {
			r := &Projectssiteinspections{InspectionID: p.InspectionID}
			db.First(r)
			fmt.Println("ProjectSiteInspections not 0")
			if err := db.Model(r).Updates(p).Error; err != nil {
				fmt.Println(err)
				return err
			}

			return nil
		}

	*/

	fmt.Println("record created")
	// Generate new primary key.

	if err := db.Save(p).Error; err != nil {
		fmt.Println("orm/Clients.go/createOrUpdateSiteInspection theres an error where it calles the Create() function")
		fmt.Println(err)
		return err
	}

	fmt.Println("orm/Clients.go/createOrUpdateSiteInspection all one big jebait")
	return nil
}
