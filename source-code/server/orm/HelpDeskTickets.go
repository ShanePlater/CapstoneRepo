package orm

import (
	"errors"
	"server/utils"

	"github.com/jinzhu/gorm"
)

// CreateHelpDeskTicket creates a new help desk ticket.
func CreateHelpDeskTicket(t *Helpdesktickets, db *gorm.DB) error {
	var last Helpdesktickets

	// Set status.
	t.Status = "1"

	if utils.Atoi(last.TicketID) != 0 {
		// Generate new primary key.
		for {
			db.Last(&last)
			t.TicketID = utils.Itoa(utils.Atoi(last.TicketID) + 1)
			check := Helpdesktickets{TicketID: t.TicketID}

			// Break for loop if primary key is available.
			if db.NewRecord(check) {
				break
			}
		}
	} else {
		// In case that table is empty.
		t.TicketID = "1"
	}

	if err := db.Create(t).Error; err != nil {
		return err
	}

	// Check if News is created.
	if db.NewRecord(*t) {
		return errors.New("fail to create help desktop ticket")
	}

	return nil

}
