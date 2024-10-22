package models

import (
	"reflect"
	"server/types"
	"server/utils"
	"sort"
	"strings"
)

type byAuthorizedDate []interface{}

// Helper functions for sortDocResources.
func (a byAuthorizedDate) Len() int      { return len(a) }
func (a byAuthorizedDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byAuthorizedDate) Less(i, j int) bool {
	return reflect.ValueOf(a[i]).FieldByName("AuthorizedDate").String() > reflect.ValueOf(a[j]).FieldByName("AuthorizedDate").String()
}

// GetRecentCompanyDocumentResources return most recent document, where length is number of document records to be returned.
func (c *Context) GetRecentCompanyDocumentResources(id, length int) []interface{} {
	res := &copy{context: c}
	res.reserveInt = append(res.reserveInt, id)

	// Loop all records in CompanyDocumentResources Table.
	c.GetCompanyDocumentResourcesTable().Range(res.rangeCompanyDocumentResources)

	// Sort interface slice by date/time in order of latest to oldest.
	sort.Sort(byAuthorizedDate(res.intf))

	// -1 is meant to ignore number of document records to be returned (return as more as possible).
	if length == -1 {
		return res.intf
	}

	// Return slice as specific number of document records is required.
	return res.intf[:length]
}

func (a *copy) rangeCompanyDocumentResources(key, value interface{}) bool {
	categoryID := reflect.ValueOf(value).FieldByName("CategoryID").String()

	// Skip if categoryID is not equal to required id.
	// -1 is meant to ignore categoryID (loop all records).
	if a.reserveInt[0] != -1 && categoryID != utils.Itoa(a.reserveInt[0]) {
		return true
	}

	// Get relevant record in CompanyDocumentResourceCategories Table.
	c, ok := a.context.GetCompanyDocumentResourceCategoriesTable().Load(categoryID)
	if !ok {
		return true
	}

	// Pass value to interface in copy.
	a.intf = append(a.intf, types.Resource{
		FileID:         reflect.ValueOf(value).FieldByName("FileName").String(),
		FileName:       reflect.ValueOf(value).FieldByName("FileFriendlyName").String(),
		FileRevision:   reflect.ValueOf(value).FieldByName("FileRevision").String(),
		AuthorizedBy:   reflect.ValueOf(value).FieldByName("AuthorizedBy").String(),
		AuthorizedDate: reflect.ValueOf(value).FieldByName("AuthorizedDate").String(),
		CategoryID:     utils.Atoi(categoryID),
		URL:            "/static/" + reflect.ValueOf(c).FieldByName("CategoryFolderPath").String() + "/" + reflect.ValueOf(value).FieldByName("FileName").String(),
	})

	// Return true for keep looping through map. Otherwise, loop exit.
	return true
}

// SearchResources return Resources which achieve the given requirements.
func (c *Context) SearchResources(keyword string) interface{} {
	// Set keyword for search.
	res := &copy{
		context:       c,
		reserveString: []string{keyword}}

	// Obtain records which contains the keyword.
	c.GetCompanyDocumentResourcesTable().Range(res.rangeResourcesSearch)

	// Sort records in alphabetical order.
	sort.Sort(byID(res.intf))

	return res.intf
}

// rangeResourcesSearch pass records which achieve requirements.
func (a *copy) rangeResourcesSearch(key, value interface{}) bool {
	categoryID := reflect.ValueOf(value).FieldByName("CategoryID").String()

	// Get relevant record in CompanyDocumentResourceCategories Table.
	c, ok := a.context.GetCompanyDocumentResourceCategoriesTable().Load(categoryID)
	if !ok {
		return true
	}

	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		if strings.Contains(reflect.ValueOf(value).Field(i).String(), a.reserveString[0]) {
			p := types.Resource{
				FileID:         reflect.ValueOf(value).FieldByName("FileName").String(),
				FileName:       reflect.ValueOf(value).FieldByName("FileFriendlyName").String(),
				FileRevision:   reflect.ValueOf(value).FieldByName("FileRevision").String(),
				AuthorizedBy:   reflect.ValueOf(value).FieldByName("AuthorizedBy").String(),
				AuthorizedDate: reflect.ValueOf(value).FieldByName("AuthorizedDate").String(),
				CategoryID:     utils.Atoi(categoryID),
				URL:            "/static/" + reflect.ValueOf(c).FieldByName("CategoryFolderPath").String() + "/" + reflect.ValueOf(value).FieldByName("FileName").String(),
			}

			// Pass value to interface in copy.
			a.intf = append(a.intf, p)

			return true
		}
	}
	return true
}
