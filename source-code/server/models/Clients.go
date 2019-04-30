package models

import (
	"reflect"
	"server/types"
	"sort"
	"strings"
)

// SearchClients return Clients which achieve the given requirements.
func (c *Context) SearchClients(keyword string) interface{} {
	// Set keyword for search.
	res := &copy{reserveString: []string{keyword}}

	// Obtain records which contains the keyword.
	c.GetClientsTable().Range(res.rangeClientsSearch)

	// Sort records in order latest to oldest.
	sort.Sort(byID(res.intf))

	return res.intf
}

// rangeClientsSearch pass records which achieve requirements.
func (a *copy) rangeClientsSearch(key, value interface{}) bool {
	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		if strings.Contains(reflect.ValueOf(value).Field(i).String(), a.reserveString[0]) {
			p := types.Client{
				ID:          reflect.ValueOf(value).FieldByName("ClientID").String(),
				Name:        reflect.ValueOf(value).FieldByName("ClientName").String(),
				PhoneNumber: reflect.ValueOf(value).FieldByName("MobileNumber").String(),
				FaxNumber:   reflect.ValueOf(value).FieldByName("FaxNumber").String(),
				Address:     reflect.ValueOf(value).FieldByName("StreetAddress").String(),
			}

			// Append address information
			p.Address = p.Address + ", " + reflect.ValueOf(value).FieldByName("StreetSuburb").String()
			p.Address = p.Address + ", " + reflect.ValueOf(value).FieldByName("StreetState").String()
			p.Address = p.Address + ", " + reflect.ValueOf(value).FieldByName("StreetPostcode").String()

			// Pass value to interface in copy.
			a.intf = append(a.intf, p)

			return true
		}
	}
	return true
}
