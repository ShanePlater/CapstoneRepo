package orm

import (
	"errors"
	"server/utils"

	"github.com/jinzhu/gorm"
)
func CreateProject(p *Projects, db *gorm.DB) error {
	if utils.Atoi(last.ID) != 0{
		// generate new primary key
		for{
			db.Last(&last)
			p.ID = utils.Itoa(utils.Atoi(last.ID) + 1)
			check := Projects{ID: p.ID}
			// break for loop if primary key is available
			if db.NewRecord(check){
				break
			}
		}
	}else{
		// project table is empty
		p.ID = "1"
	}

	if err := db.Create(p).Error; err != nil{
		return err
	}

	if db.NewRecord(*p){
		return errors.New("fail to create new project")
	}

	return nil
}