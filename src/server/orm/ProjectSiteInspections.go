package orm

import (
	"errors"
	"fmt"
	"server/utils"

	"github.com/jinzhu/gorm"
)

//CreateOrUpdateProjectSiteInspections facilitating creation of new SiteInspections
func CreateOrUpdateProjectSiteInspections(p *Projectssiteinspections, db *gorm.DB) error {
	fmt.Println("orm/Clients.go/createOrUpdateClientss to create Client in orm")
	// Update record.
	if p.InspectionID != "0" {
		r := &Projectssiteinspections{InspectionID: p.InspectionID}
		db.First(r)
		fmt.Println("ProjectSiteInspections not 0")
		if err := db.Model(r).Updates(p).Error; err != nil {
			fmt.Println("err")
			return err
		}

		return nil
	}

	fmt.Println("survived 1st if")

	// Create record.
	last := &Projectssiteinspections{}
	fmt.Println("record created")
	// Generate new primary key.
	var primkeyloopcount int
	for {
		db.Last(last)
		p.InspectionID = utils.Itoa(utils.Atoi(last.InspectionID) + 1)
		check := Projectssiteinspections{InspectionID: p.InspectionID}
		primkeyloopcount++

		fmt.Println(primkeyloopcount)

		// Break for loop if primary key is available.
		if db.NewRecord(check) {
			fmt.Println("primary key made")
			break
		}
	}

	fmt.Println("primary key generated")

	if err := db.Create(p).Error; err != nil {
		fmt.Println("orm/Clients.go/createOrUpdateSiteInspection theres an error where it calles the Create() function")
		return err
	}

	// Check if News is created successfully.
	if db.NewRecord(p) {
		return errors.New("failed to create new Client")
	}

	return nil
}
