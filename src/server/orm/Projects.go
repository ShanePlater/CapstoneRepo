package orm

import (
	"errors"
	"fmt"
	"server/utils"
	"strings"

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
		p.ProjectNumber = last.ProjectNumber

		//This is very messy as for the primary keys L&R uses a character instead of an int
		//it means when we try and add 1 to the primary key, it breaks the system as its a char

		//grab the last 5 chars in the string (which will always be numbers)
		lastproj := utils.Atoi(string(p.ProjectNumber[len(p.ProjectNumber)-5:]))
		//create the new primary key by adding 1 to the old number
		newkey := lastproj + 1
		//replace the old 4 digits with the new four digits.
		strings.Replace(p.ProjectNumber, utils.Itoa(lastproj), utils.Itoa(newkey), -1)
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
	fmt.Println("orm/Projects.go/createOrUpdateProjects End of ORM")
	return nil
}
