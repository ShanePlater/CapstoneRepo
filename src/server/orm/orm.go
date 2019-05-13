package orm

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// MSSQL object relational mapping driver.
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

// OpenDB initializes a new db connection.
func OpenDB(mssql string) *gorm.DB {
	db, err := gorm.Open("mssql", mssql)
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}

	// Disable table name's pluralization globally
	db.SingularTable(true)

	return db
}
