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

// CreateAttachedResource creates the file record.
func CreateAttachedResource(f *Companydocumentresources, db *gorm.DB) error {
	if !db.NewRecord(f.FileName) {
		return errors.New("this file already exists")
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

// DeleteAttachedResource deletes the file record.
func DeleteAttachedResource(fileName *string, db *gorm.DB) error {
	f := Companydocumentresources{FileName: *fileName}
	if err := db.Delete(&f).Error; err != nil {
		return err
	}
	return nil
}
