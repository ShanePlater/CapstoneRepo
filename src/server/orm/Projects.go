package orm

import (
	"errors"
	"fmt"
	"server/utils"

	"github.com/jinzhu/gorm"
)

//CreateOrUpdateProjects creates project
func CreateOrUpdateProjects(p *Projects, db *gorm.DB) error {
	fmt.Println("orm/Projects.go/createOrUpdateProjects to create project in orm")
	// Update record.
	if p.ProjectNumber != "0" {
		r := &Projects{ProjectNumber: p.ProjectNumber}
		db.First(r)

		if err := db.Model(r).Updates(p).Error; err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}

	// Create record.
	last := &Projects{}

	// Generate new primary key.
	for {
		db.Last(last)
		p.ProjectNumber = utils.Itoa(utils.Atoi(last.ProjectNumber) + 1)
		check := Projects{ProjectNumber: p.ProjectNumber}

		// Break for loop if primary key is available.
		if db.NewRecord(check) {
			break
		}
	}

	if err := db.Create(p).Error; err != nil {
		fmt.Println("orm/Projects.go/createOrUpdateProjects theres an error where it calles the Create() function")
		return err
	}

	// Check if News is created successfully.
	if db.NewRecord(p) {
		return errors.New("failed to create new project")
	}

	return nil
}
