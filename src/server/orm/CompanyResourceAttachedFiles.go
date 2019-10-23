package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// UpdateAttachedResource updates the attached file information.
func UpdateAttachedResource(f *Companydocumentresources, db *gorm.DB) error {
	r := &Companydocumentresources{FileName: f.FileName}
	db.First(r)

	if err := db.Model(r).Updates(f).Error; err != nil {
		return err
	}

	return nil
}

// CreateOrUpdateAttachedResource creates the file record.
func CreateOrUpdateAttachedResource(f *Companydocumentresources, db *gorm.DB) error {

	if f.FileRevision != "A" && f.FileRevision != "a" {
		r := &Companydocumentresources{FileName: f.FileName}
		db.First(r)
		if err := db.Model(r).Updates(f).Error; err != nil {
			return err
		}
	} else {
		if err := db.Create(f).Error; err != nil {
			return err
		}
	}
	// Check if new record is stored successfully.
	if db.NewRecord(f) {
		return errors.New("fail to create new file entity")
	}

	return nil
}

// DeleteAttachedResource deletes the file record.
func DeleteAttachedResource(fileName *string, db *gorm.DB) error {
	f := Companydocumentresources{FileName: *fileName}
	if err := db.Delete(&f).Error; err != nil {
		return err
	}
	return nil
}
