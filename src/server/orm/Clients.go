package orm

import (
	"errors"
	"fmt"
	"server/utils"

	"github.com/jinzhu/gorm"
)

//CreateOrUpdateClients facilitating creation of new clients
func CreateOrUpdateClients(p *Clients, db *gorm.DB) error {
	fmt.Println("orm/Clients.go/createOrUpdateClientss to create Client in orm")
	// Update record.
	if p.ClientID != "0" {
		r := &Clients{ClientID: p.ClientID}
		db.First(r)
		fmt.Println("clientID not 0")
		if err := db.Model(r).Updates(p).Error; err != nil {
			fmt.Println("err")
			return err
		}

		return nil
	}

	fmt.Println("survived 1st if")

	// Create record.
	last := &Clients{}
	fmt.Println("record created")
	// Generate new primary key.
	for {
		db.Last(&last)
		p.ClientID = utils.Itoa(utils.Atoi(last.ClientID) + 1)

		/*
			check := Clients{ClientID: p.ClientID}
				primkeyloopcount++
				if primkeyloopcount < 5 {
					fmt.Println(last.ClientID + "check vs" + p.ClientID)
					fmt.Println(db.NewRecord(check))
				}

				// Break for loop if primary key is available.
				if db.NewRecord(check) {
					fmt.Println("primary key made")
					break
				}
		*/
		break
	}

	fmt.Println("primary key generated")

	if err := db.Create(p).Error; err != nil {
		fmt.Println("orm/Clientss.go/createOrUpdateClientss theres an error where it calles the Create() function")
		return err
	}

	// Check if News is created successfully.
	if db.NewRecord(p) {
		return errors.New("failed to create new Client")
	}

	return nil
}
