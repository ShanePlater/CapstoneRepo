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
		fmt.Println("orm/Projects.go/createOrUpdateProjects projnum != 0")
		r := &Projects{ProjectNumber: p.ProjectNumber}
		fmt.Println("orm/Projects.go/createOrUpdateProjects proj num of updated proj is:" + p.ProjectNumber)
		db.First(r)
		fmt.Println("orm/Projects.go/createOrUpdateProjects proj num of updated proj is:" + p.ProjectNumber + "DB FOUND")

		if err := db.Model(r).Updates(p).Error; err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}

	// Create record.
	last := &Projects{}

	// Generate new primary key.

	db.Last(last)
	p.ProjectNumber = last.ProjectNumber

	//This is very messy as for the primary keys L&R uses a character instead of an int
	//it means when we try and add 1 to the primary key, it breaks the system as its a char
	//WXXXXX
	//grab the last 5 chars in the string (which will always be numbers)
	lastproj := utils.Atoi(string(p.ProjectNumber[len(p.ProjectNumber)-5:]))
	fmt.Println("orm/Projects.go/createOrUpdateProjects lastproj is: " + utils.Itoa(lastproj))
	//create the new primary key by adding 1 to the old number
	newdigit := lastproj + 1
	fmt.Println("orm/Projects.go/createOrUpdateProjects newdigit is: " + utils.Itoa(newdigit))
	newkey := "W" + utils.Itoa(newdigit)
	p.ProjectNumber = newkey
	fmt.Println("orm/Projects.go/createOrUpdateProjects new project number is is: " + p.ProjectNumber)

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
