package orm

import (
	"errors"
	"server/utils"

	"github.com/jinzhu/gorm"
)

// UpdateAttachedFile updates the attached file information.
func UpdateAttachedFile(f *Helpdeskticketattachedfiles, db *gorm.DB) error {
	r := &Helpdeskticketattachedfiles{FileID: f.FileID}
	db.First(r)

	if err := db.Model(r).Updates(f).Error; err != nil {
		return err
	}

	return nil
}

// CreateAttachedFile creates the file record.
func CreateAttachedFile(f *Helpdeskticketattachedfiles, db *gorm.DB) error {
	var last Helpdeskticketattachedfiles

	db.Last(&last)
	if utils.Atoi(last.FileID) != 0 {
		// Generate new primary key.
		for {
			db.Last(&last)
			f.FileID = utils.Itoa(utils.Atoi(last.FileID) + 1)
			check := Helpdeskticketattachedfiles{FileID: f.FileID}

			// Break for loop if primary key is available.
			if db.NewRecord(check) {
				break
			}
		}
	} else {
		// In case that table is empty.
		f.FileID = "1"
	}

	if err := db.Create(f).Error; err != nil {
		return err
	}

	// Check if new record is stored successfully.
	if db.NewRecord(f) {
		return errors.New("fail to create new file entity")
	}

	return nil
}

// DeleteAttachedFile deletes the file record.
func DeleteAttachedFile(id *int, db *gorm.DB) error {
	f := Helpdeskticketattachedfiles{FileID: utils.Itoa(*id)}
	if err := db.Delete(&f).Error; err != nil {
		return err
	}
	return nil
}
