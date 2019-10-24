package models

import (
	"reflect"
	"server/utils"
	"strings"
)

/*Give this method the office code and division name and it will return the code.
*This is important as the front end sends the Division name to the server, but the database only takes division codes for projects.
* So this method is needed to convert user entries for projects to a DB compatible format.
*/
func (c *Context) locateDivisionCode(officeCode, divisionName string) int {
	res := &copy{reserveString: []string{officeCode, divisionName}, reserveInt: []int{-1}}
	c.GetDivisionsTable().Range(res.rangeDivisionsForDivisionCode)
	return res.reserveInt[0]
}

func (a *copy) rangeDivisionsForDivisionCode(key, value interface{}) bool {
	if a.reserveString[0] != "" && !strings.Contains(reflect.ValueOf(value).FieldByName("OfficeCode").String(), a.reserveString[0]) {
		return true
	}
	if a.reserveString[1] != "" && reflect.ValueOf(value).FieldByName("DivisionName").String() != a.reserveString[1] {
		return true
	}
	a.reserveInt = []int{utils.Atoi(reflect.ValueOf(key).String())}
	return false
}
