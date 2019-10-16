package orm

import (
	"errors"
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

	if err := db.Create(p).Error; err != nil {
		fmt.Println("orm/Clients.go/createOrUpdateSiteInspection theres an error where it calles the Create() function")
		fmt.Println(err)
		return err
	}

	// Check if News is created successfully.
	if db.NewRecord(p) {
		return errors.New("This always returns false, we don't know why")
	}

	return nil
}
