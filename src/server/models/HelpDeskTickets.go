package models

import (
	"reflect"
	"server/orm"
	"server/types"
	"server/utils"
	"sort"
)

// CreateHelpDeskTicket creates a new help desk ticket.
func (c *Context) CreateHelpDeskTicket(data *types.HelpDeskTicket) error {
	r := orm.Helpdesktickets{
		Situation:   utils.Itoa(data.Situation),
		Username:    data.Username,
		Computer:    data.Computer,
		Description: data.Description,
	}

	if err := orm.CreateHelpDeskTicket(&r, c.db); err != nil {
		return err
	}

	// Check if cache is enabled.
	if c.config.IsCache() {
		c.helpDeskTickets.Store(r.TicketID, r)
	}

	// Update data.
	*data = types.HelpDeskTicket{
		TicketID:    utils.Atoi(r.TicketID),
		Status:      utils.Atoi(r.Status),
		Username:    r.Username,
		Computer:    r.Computer,
		Description: r.Description,
	}

	return nil
}

type byTicketID []interface{}

// Helper functions for sortNews.
func (a byTicketID) Len() int      { return len(a) }
func (a byTicketID) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byTicketID) Less(i, j int) bool {
	return reflect.ValueOf(a[i]).FieldByName("TicketID").String() > reflect.ValueOf(a[j]).FieldByName("TicketID").String()
}

// GetHelpDeskTickets return all records related to Help Desk Tickets.
func (c *Context) GetHelpDeskTickets() interface{} {
	// Set given time period.
	res := &copy{}

	// Obtain records in time period
	c.GetHelpDeskTicketsTable().Range(res.rangeHelpDeskTickets)

	// Sort records in order of largest to smallest.
	sort.Sort(byTicketID(res.intf))

	return res.intf

}

// rangeHelpDeskTickets pass all records of HelpDeskTickets table to copy.
func (a *copy) rangeHelpDeskTickets(key, value interface{}) bool {
	files := &copy{}

	// Get help desk attached files by TicketID.
	files.reserveString = append(files.reserveString, reflect.ValueOf(value).FieldByName("TicketID").String())
	a.context.GetHelpDeskTicketAttachedFilesTable().Range(files.rangeHelpDeskTicketAttachedFiles)

	// Pass value to interface in copy.
	a.intf = append(a.intf, types.HelpDeskTicketInterface{
		TicketID:    utils.Atoi(reflect.ValueOf(value).FieldByName("TicketID").String()),
		Status:      utils.Atoi(reflect.ValueOf(value).FieldByName("Status").String()),
		Situation:   utils.Atoi(reflect.ValueOf(value).FieldByName("Situation").String()),
		Username:    reflect.ValueOf(value).FieldByName("Username").String(),
		Computer:    reflect.ValueOf(value).FieldByName("Computer").String(),
		Description: reflect.ValueOf(value).FieldByName("Description").String(),
		Files:       files.intf,
	})

	// Return true for keep looping through map. Otherwise, loop exit.
	return true
}
