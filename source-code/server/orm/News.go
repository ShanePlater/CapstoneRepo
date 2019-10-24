package orm

import (
	"errors"
	"server/utils"

	"github.com/jinzhu/gorm"
)

// CreateOrUpdateNews update a News record. If ID is not defined, it will create a new News record.
func CreateOrUpdateNews(n *News, db *gorm.DB) error {
	// Update record.
	if n.ArticleID != "0" {
		r := &News{ArticleID: n.ArticleID}
		db.First(r)

		if err := db.Model(r).Updates(n).Error; err != nil {
			return err
		}

		return nil
	}

	// Create record.
	last := &News{}

	// Generate new primary key.
	for {
		db.Last(last)
		n.ArticleID = utils.Itoa(utils.Atoi(last.ArticleID) + 1)
		check := News{ArticleID: n.ArticleID}

		// Break for loop if primary key is available.
		if db.NewRecord(check) {
			break
		}
	}

	if err := db.Create(n).Error; err != nil {
		return err
	}

	// Check if News is created successfully.
	if db.NewRecord(n) {
		return errors.New("fail to create news")
	}

	return nil
}

// RemoveNews remove a given News record.
func RemoveNews(id *int, db *gorm.DB) error {
	r := &News{ArticleID: utils.Itoa(*id)}

	if err := db.Delete(r).Error; err != nil {
		return err
	}

	return nil
}
