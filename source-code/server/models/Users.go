package models

import (
	"reflect"
	"server/types"
	"server/utils"
	"sort"
	"strings"
)

// SearchUsers return Users which achieve the given requirements.
func (c *Context) SearchUsers(keyword string) interface{} {
	// Set keyword for search.
	res := &copy{reserveString: []string{keyword}}

	// Obtain records which contains the keyword.
	c.GetUsersTable().Range(res.rangeUsersSearch)

	// Sort records in alphabetical order.
	sort.Sort(byID(res.intf))

	return res.intf
}

// rangeUsersSearch pass records which achieve requirements.
func (a *copy) rangeUsersSearch(key, value interface{}) bool {
	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		if strings.Contains(reflect.ValueOf(value).Field(i).String(), a.reserveString[0]) {
			p := types.User{
				ID:             reflect.ValueOf(value).FieldByName("Username").String(),
				Name:           reflect.ValueOf(value).FieldByName("FirstName").String() + " " + reflect.ValueOf(value).FieldByName("LastName").String(),
				PhoneNumber:    reflect.ValueOf(value).FieldByName("MobileNumber").String(),
				PhoneExtension: reflect.ValueOf(value).FieldByName("PhoneExtension").String(),
				Email:          reflect.ValueOf(value).FieldByName("EmailAddress").String(),
				Division:       reflect.ValueOf(value).FieldByName("DivisionCode").String(),
			}

			// Pass value to interface in copy.
			a.intf = append(a.intf, p)

			return true
		}
	}
	return true
}

func (c *Context) GetUsersByDivisionCode(data *types.GetDivisionsJSON) interface{} {
	divisionCode := c.locateDivisionCode(data.Office, data.Division)
	if divisionCode != -1 {
		res := &copy{reserveInt: []int{divisionCode}}
		c.GetUsersTable().Range(res.rangeUsersByDivisionCode)
		return res.intf
	}
	return nil
}

func (a *copy) rangeUsersByDivisionCode(key, value interface{}) bool {
	v := reflect.ValueOf(value)
	if v.FieldByName("DivisionCode").String() == utils.Itoa(a.reserveInt[0]) {
		a.intf = append(a.intf, types.Division{
			EmailAddress:   v.FieldByName("EmailAddress").String(),
			Name:           v.FieldByName("FirstName").String() + " " + v.FieldByName("LastName").String(),
			MobileNumber:   v.FieldByName("MobileNumber").String(),
			PhoneExtension: v.FieldByName("PhoneExtension").String(),
			PhotoURL:       "/static/" + v.FieldByName("PhotoURL").String(),
			PositionTitle:  v.FieldByName("PositionTitle").String(),
			Username:       v.FieldByName("Username").String(),
		})
	}
	return true
}
