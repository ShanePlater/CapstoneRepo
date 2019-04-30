package models

import (
	"reflect"
	"server/utils"
	"strings"
)

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
