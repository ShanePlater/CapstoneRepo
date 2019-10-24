package models

import (
	"errors"
	"reflect"
	"server/orm"
	"server/types"
	"server/utils"
)

// UpdateAttachedFile updates the attached file information.
func (c *Context) UpdateAttachedFile(data *types.File, tktID int) error {
	r := orm.Helpdeskticketattachedfiles{
		FileID:   utils.Itoa(data.FileID),
		TicketID: utils.Itoa(tktID),
		FileName: data.FileName,
	}

	// Call function in orm to update attached file record stored in database.
	if err := orm.UpdateAttachedFile(&r, c.db); err != nil {
		return err
	}

	// Check if cache is enabled.
	if c.config.IsCache() {
		c.helpDeskTicketAttachedFiles.Store(r.FileID, r)
	}

	// Update data.
	*data = types.File{
		FileID:   utils.Atoi(r.FileID),
		FileName: r.FileName,
		URL:      "/pub/HelpDesk/" + r.FileID,
	}

	return nil
}

// rangeHelpDeskTicketAttachedFiles pass all records with given TicketID to copy.
func (a *copy) rangeHelpDeskTicketAttachedFiles(key, value interface{}) bool {
	if a.reserveString[0] != reflect.ValueOf(value).FieldByName("TicketID").String() {
		return true
	}
	// Add record to the result.
	a.intf = append(a.intf, types.File{
		FileID:   utils.Atoi(reflect.ValueOf(value).FieldByName("FileID").String()),
		FileName: reflect.ValueOf(value).FieldByName("FileName").String(),
		URL:      "/pub/HelpDesk/" + reflect.ValueOf(value).FieldByName("FileID").String(),
	})

	// Return true for keep looping through map. Otherwise, loop exit.
	return true
}

// CreateAttachedFile updates the attached file information.
func (c *Context) CreateAttachedFile(data *types.File) error {
	r := orm.Helpdeskticketattachedfiles{FileName: data.FileName}

	if err := orm.CreateAttachedFile(&r, c.db); err != nil {
		return err
	}

	// Check if cache is enabled.
	if c.config.IsCache() {
		c.helpDeskTicketAttachedFiles.Store(r.FileID, r)
	}

	*data = types.File{
		FileID:   utils.Atoi(r.FileID),
		FileName: r.FileName,
		URL:      "/pub/HelpDesk/" + r.FileID,
	}

	return nil
}

// DeleteAttachedFile deletes the file record.
func (c *Context) DeleteAttachedFile(id int) error {
	if err := orm.DeleteAttachedFile(&id, c.db); err != nil {
		return err
	}

	// Check if cache is enabled.
	if !c.config.IsCache() {
		return nil
	}

	c.helpDeskTicketAttachedFiles.Delete(id)

	if _, ok := c.helpDeskTicketAttachedFiles.Load(id); ok {
		return errors.New("fail to remove attached file")
	}

	return nil
}
