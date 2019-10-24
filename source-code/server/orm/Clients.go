package orm

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

//CreateOrUpdateClients facilitating creation of new clients
func CreateOrUpdateClients(p *Clients, db *gorm.DB, ClientID string) error {
	fmt.Println("orm/Clients.go/createOrUpdateClientss to create Client in orm")
	// Update record.
	if ClientID != "0" {
		r := &Clients{ClientID: ClientID}
		db.First(r)
		fmt.Println("clientID not 0")
		if err := db.Model(r).Update(p).Error; err != nil {
			fmt.Println("err")
			return err
		}

		return nil
	}

	fmt.Println("survived 1st if")

	// Create record.
	last := &Clients{}
	fmt.Println("record created")

	for {
		db.Last(&last)
		//p.ClientID = utils.Itoa(utils.Atoi(last.ClientID) + 1)

		//check := Clients{ClientID: p.ClientID}

		break
	}

	fmt.Println("primary key generated")

	if err := db.Save(p).Error; err != nil {
		fmt.Println("orm/Clientss.go/createOrUpdateClientss theres an error where it calles the Create() function")
		return err
	}

	// Check if News is created successfully.
	if db.NewRecord(p) {
		return errors.New("failed to create new Client")
	}
	return nil
}

func setIdentityInsert(scope *gorm.Scope) {
	if scope.Dialect().GetName() == "mssql" {
		for _, field := range scope.PrimaryFields() {
			if _, ok := field.TagSettings["AUTO_INCREMENT"]; ok && !field.IsBlank {
				scope.NewDB().Exec(fmt.Sprintf("SET IDENTITY_INSERT %v ON", scope.TableName()))
				scope.InstanceSet("mssql:identity_insert_on", true)
			}
		}
	}
}

func turnOffIdentityInsert(scope *gorm.Scope) {
	if scope.Dialect().GetName() == "mssql" {
		if _, ok := scope.InstanceGet("mssql:identity_insert_on"); ok {
			scope.NewDB().Exec(fmt.Sprintf("SET IDENTITY_INSERT %v OFF", scope.TableName()))
		}
	}
}
