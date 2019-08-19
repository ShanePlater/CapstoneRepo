package models

import (
	"fmt"
	"reflect"
	"server/orm"
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

//CreateOrUpdateClients create client class and upload to database
func (c *Context) CreateOrUpdateClients(data *types.Client2) error {
	fmt.Println("models/ CreatingOrUpdateClients")

	//Get division code, used Brisbane as the default office as we dont have proper AD sync to take it from user yet
	//data.Division = utils.Itoa(c.locateDivisionCode("BNE", data.Division))

	//current hard coded values: clientID, Client Office Code
	//Project Director & Manager --> The DB has 0 for many of these, after that it is the domain login, implemented in a later sprint.
	r := orm.Clients{
		ClientID:              "0",
		ClientOfficeCode:      "BNE",
		ClientName:            data.ClientName,
		ClientABNNumber:       data.ClientABNNumber,
		ClientACNNumber:       data.ClientACNNumber,
		ClientTypeCode:        data.ClientTypeCode,
		FirstName:             data.FirstName,
		LastName:              data.LastName,
		ClientLocationCode:    "1",
		StreetAddress:         data.StreetAddress,
		StreetSuburb:          data.StreetSuburb,
		StreetState:           data.StreetState,
		StreetPostcode:        data.StreetPostcode,
		PostalAddress:         data.PostalAddress,
		PostalAddressSuburb:   data.PostalAddressSuburb,
		PostalAddressState:    data.PostalAddressState,
		PostalAddressPostcode: data.PostalAddressPostcode,
		PhoneNumber:           data.PhoneNumber,
		FaxNumber:             data.FaxNumber,
		MobileNumber:          data.MobileNumber,
		EMailAddress:          data.EMailAddress,
	}

	if err := orm.CreateOrUpdateClients(&r, c.db); err != nil {
		fmt.Println("Models/ Error in Clients.go")
		return err
	}

	// Check if cache is enabled.
	if c.config.IsCache() {
		c.clients.Store(r.ClientID, r)
	}

	// Update data.
	*data = types.Client2{
		ClientID:              r.ClientID,
		ClientOfficeCode:      r.ClientOfficeCode,
		ClientName:            r.ClientName,
		ClientABNNumber:       r.ClientABNNumber,
		ClientACNNumber:       r.ClientACNNumber,
		ClientTypeCode:        r.ClientTypeCode,
		FirstName:             r.FirstName,
		LastName:              r.LastName,
		ClientLocationCode:    r.ClientLocationCode,
		StreetAddress:         r.StreetAddress,
		StreetSuburb:          r.StreetSuburb,
		StreetState:           r.StreetState,
		StreetPostcode:        r.StreetPostcode,
		PostalAddress:         r.PostalAddress,
		PostalAddressSuburb:   r.PostalAddressSuburb,
		PostalAddressState:    r.PostalAddressState,
		PostalAddressPostcode: r.PostalAddressPostcode,
		PhoneNumber:           r.PhoneNumber,
		FaxNumber:             r.FaxNumber,
		MobileNumber:          r.MobileNumber,
		EMailAddress:          r.EMailAddress,
	}

	return nil
}
