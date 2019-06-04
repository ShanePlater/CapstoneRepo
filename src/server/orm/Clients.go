package orm

import (
	"errors"
	"fmt"
	"server/utils"

	"github.com/jinzhu/gorm"
)

// creates Client
func CreateOrUpdateClients(p *Clients, db *gorm.DB) error {
	fmt.Println("orm/Clients.go/createOrUpdateClientss to create Client in orm")
	// Update record.
	if p.ClientID != "0" {
		r := &Clients{ClientID: p.ClientID}
		db.First(r)

		if err := db.Model(r).Updates(p).Error; err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}

	// Create record.
	last := &Clients{}

	// Generate new primary key.
	for {
		db.Last(last)
		p.ClientID = utils.Itoa(utils.Atoi(last.ClientID) + 1)
		check := Clients{ClientID: p.ClientID}

		// Break for loop if primary key is available.
		if db.NewRecord(check) {
			break
		}
	}

	if err := db.Create(p).Error; err != nil {
		fmt.Println("orm/Clientss.go/createOrUpdateClientss theres an error where it calles the Create() function")
		return err
	}

	// Check if News is created successfully.
	if db.NewRecord(p) {
		return errors.New("Failed to create new Client.")
	}

	return nil
}
