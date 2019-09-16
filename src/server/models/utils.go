package models

import (
	"bytes"
	"html/template"
	"net/smtp"
	"reflect"
	"regexp"
	"server/types"
	"sort"
	"strings"
	"sync"
)

type byName []interface{}
type byDate []interface{}
type byID []interface{}

// Helper functions for sorting by Name.
func (a byName) Len() int      { return len(a) }
func (a byName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool {
	return reflect.ValueOf(a[i]).FieldByName("Name").String() < reflect.ValueOf(a[j]).FieldByName("Name").String()
}

// Helper functions for sorting by Date.
func (a byDate) Len() int      { return len(a) }
func (a byDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byDate) Less(i, j int) bool {
	return reflect.ValueOf(a[i]).FieldByName("Date").String() > reflect.ValueOf(a[j]).FieldByName("Date").String()
}

// Helper functions for sorting by ID.
func (a byID) Len() int      { return len(a) }
func (a byID) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool {
	return reflect.ValueOf(a[i]).FieldByName("ID").String() < reflect.ValueOf(a[j]).FieldByName("ID").String()
}

// inTimePeriod checks whether a given time is in a given time period.
func inTimePeriod(start, end, check string) bool {
	// Cut "2006-01-02T15:04:05Z07:00" or "2006-01-02 00:00:00.000" to "2006-01-02".
	return start[:10] <= check[:10] && check[:10] <= end[:10]
}

// GetOptions return records relate to
// ProjectLocations table, ProjectTypes table, ProjectStatuss table,
// Divisions table, and Offices table.
func (c *Context) GetOptions(name string) interface{} {
	// Set given time period.
	res := &copy{}

	// Obtain records.
	switch name {
	case "getOptionLocations":
		res.reserveString = append(res.reserveString, "ProjectLocationCode", "ProjectLocationName")
		c.GetProjectLocationsTable().Range(res.rangeOptions)
	case "getOptionTypes":
		res.reserveString = append(res.reserveString, "ProjectTypeCode", "ProjectTypeName")
		c.GetProjectTypesTable().Range(res.rangeOptions)
	case "getOptionStatuss":
		res.reserveString = append(res.reserveString, "ProjectStatusCode", "ProjectStatusName")
		c.GetProjectStatussTable().Range(res.rangeOptions)
	case "getOptionDivisions":
		res.reserveString = append(res.reserveString, "FriendlyCode", "DivisionName")
		c.GetDivisionsTable().Range(res.rangeOptions)
	case "getOptionOffices":
		res.reserveString = append(res.reserveString, "OfficeCode", "Name")
		c.GetOfficesTable().Range(res.rangeOptions)
	case "getClientTypes":
		res.reserveString = append(res.reserveString, "ClientTypeCode", "CodeDescription")
		c.GetClientTypesTable().Range(res.rangeOptions)
	case "getClientLocations":
		res.reserveString = append(res.reserveString, "ClientLocationCode", "ClientLocation")
		c.GetClientLocationsTable().Range(res.rangeOptions)
	case "getOptionCategories":
		res.reserveString = append(res.reserveString, "CategoryID", "CategoryName")
		c.GetCompanyDocumentResourceCategoriesTable().Range(res.rangeOptions)
	default:
	}

	// Sort records in alphabetical order.
	sort.Sort(byName(res.intf))

	return res.intf
}

// rangeOptions pass all records to copy.
func (a *copy) rangeOptions(key, value interface{}) bool {
	name := reflect.ValueOf(value).FieldByName(a.reserveString[1]).String()
	// Check if selected value has name in empty.
	if strings.Replace(name, " ", "", -1) == "" {
		return true
	}

	// Check if selected value has been appeared.
	// Because Division has repeated names.
	// The loop starts from third place as first and second places holds option information.
	for i := 2; i < len(a.reserveString); i++ {
		if name == a.reserveString[i] {
			return true
		}
	}

	// Pass value to interface in copy.
	a.intf = append(a.intf, types.Option{
		ID:   reflect.ValueOf(value).FieldByName(a.reserveString[0]).String(),
		Name: name,
	})

	a.reserveString = append(a.reserveString, name)
	// Return true for keep looping through map. Otherwise, loop exit.
	return true
}

// SendEmail sends the email.
func (c *Context) SendEmail(data *types.CreateSuggestionJSON, subject, t string) error {
	var buff bytes.Buffer

	// Create email template.
	tmpl, err := template.New("email").Parse(t)
	if err != nil {
		return err
	}

	// Parse email template.
	err = tmpl.Execute(&buff, *data)
	if err != nil {
		return err
	}

	// Set email content.
	msg := "From: " + c.config.Get("email.system.account") + "\r\n" +
		"To: " + c.config.Get("email.system.account") + "\r\n" +
		"MIME-Version: 1.0" + "\r\n" +
		"Content-type: text/html" + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		buff.String() + "\r\n"

	// Send email.
	err = smtp.SendMail(c.config.Get("email.smtp.url")+":"+c.config.Get("email.smtp.port"), smtp.PlainAuth("",
		c.config.Get("email.system.account"), c.config.Get("email.system.password"), c.config.Get("email.smtp.url")), c.config.Get("email.system.account"), []string{c.config.Get("email.system.account")}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

// FetchSuggestions completes the search keyword based on given keyword.
func (c *Context) FetchSuggestions(data *types.SearchJSON) interface{} {
	var wg sync.WaitGroup

	// Remove space at the beginning of the line.
	s := strings.Trim(trimUnexpectedCharactors(data.Keyword), " ")

	// Convert to regexp.
	// "^" means start of the line;
	// ".*" means match all ("." is required in Go's regexp).
	s = "^" + strings.Replace(s, " ", ".*", -1) + ".*"

	r1 := &copy{}
	r2 := &copy{}
	wg.Add(2)

	// Query matched data.
	// 1 means checking with case sensitive.
	go func() {
		r1.reserveString = c.querySuggestions(s, 1)
		wg.Done()
	}()

	// 0 means checking with case insensitive.
	go func() {
		r2.reserveString = c.querySuggestions(s, 0)
		wg.Done()
	}()

	wg.Wait()

	return append(r1.reserveString, r2.reserveString...)
}

func (c *Context) querySuggestions(keyword string, caseSensitive int) []string {
	// res stores the result of this function.
	var res []string

	// values stores the results from goroutines.
	values := make(chan []string)

	// Search CompanyDocumentResources table.
	go func(s chan []string) {
		res := &copy{reserveString: []string{keyword}, reserveInt: []int{caseSensitive}}
		c.GetCompanyDocumentResourcesTable().Range(res.rangeForSuggestions)
		s <- res.reserveString[1:]
	}(values)

	// Search Clients table.
	go func(s chan []string) {
		res := &copy{reserveString: []string{keyword}, reserveInt: []int{caseSensitive}}
		c.GetClientsTable().Range(res.rangeForSuggestions)
		s <- res.reserveString[1:]
	}(values)

	// Search Users table.
	go func(s chan []string) {
		res := &copy{reserveString: []string{keyword}, reserveInt: []int{caseSensitive}}
		c.GetUsersTable().Range(res.rangeForSuggestions)
		s <- res.reserveString[1:]
	}(values)

	// Search Projects table.
	go func(s chan []string) {
		res := &copy{reserveString: []string{keyword}, reserveInt: []int{caseSensitive}}
		c.GetProjectsTable().Range(res.rangeForSuggestions)
		s <- res.reserveString[1:]
	}(values)

	// Remove the repeated items.
	// In "a < 4", "4" is the number of goroutines.
	for a := 0; a < 4; a++ {
		// Loop the values returned from each goroutine.
		for _, v := range <-values {
			exist := false
			// Loop the items for result of this function.
			for _, i := range res {
				if i == v {
					exist = true
					break
				}
			}
			// Append if not exist.
			if !exist {
				res = append(res, v)
			}
		}
	}

	// Sort the result in alphabetical order.
	sort.Strings(res)

	return res
}

func (a *copy) rangeForSuggestions(key, value interface{}) bool {
	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		if !iSValueShouldBeSearched(reflect.TypeOf(value).Field(i).Name) {
			continue
		}
		v := reflect.ValueOf(value).Field(i).String()
		if matched, _ := regexp.MatchString(a.reserveString[0], v); matched {
			// Filter results already appear from case sensitive search.
			if a.reserveInt[0] != 0 {
				a.reserveString = append(a.reserveString, v)
			}
			continue
		}
		// 0 means checking with case insensitive.
		// 1 means checking with case sensitive.
		// Default is case sensitive.
		if a.reserveInt[0] == 0 {
			if matched, _ := regexp.MatchString(strings.ToUpper(a.reserveString[0]), strings.ToUpper(v)); matched {
				a.reserveString = append(a.reserveString, v)
			}
		}
	}
	return true
}

func iSValueShouldBeSearched(name string) bool {
	for _, v := range []string{"ProjectName", "ClientRepName", "ProjectAddress", "ProjectSuburb", "ProjectDescription", "FirstName", "LastName", "ClientName", "EMailAddress", "FileName", "FileFriendlyName"} {
		if v == name {
			return true
		}
	}
	return false
}

//GetProject get list of projects
func (c *Context) GetProject(data *types.GetByIDJSON) (interface{}, bool) {
	return c.GetProjectsTable().Load(data.ID)
}

//GetClient get list of clients and return in table format
func (c *Context) GetClient(data *types.GetByIDJSON) (interface{}, bool) {
	return c.GetClientsTable().Load(data.ID)
}

//GetSuggestionItemURL legacy code unsure of use
func (c *Context) GetSuggestionItemURL(data *types.SearchJSON) string {
	var wg sync.WaitGroup

	r := &copy{context: c, reserveString: []string{data.Keyword}}
	wg.Add(4)

	go func() {
		c.GetProjectsTable().Range(r.rangeProjectForURL)
		wg.Done()
	}()

	go func() {
		c.GetUsersTable().Range(r.rangeUserForURL)
		wg.Done()
	}()

	go func() {
		c.GetClientsTable().Range(r.rangeClientForURL)
		wg.Done()
	}()

	go func() {
		c.GetCompanyDocumentResourcesTable().Range(r.rangeDocumentForURL)
		wg.Done()
	}()

	wg.Wait()

	// First place holds keyword, and second place holds keyword.
	// Therefore the length is 2.
	if len(r.reserveString) == 2 {
		return r.reserveString[1]
	}

	return ""
}

func (a *copy) rangeProjectForURL(key, value interface{}) bool {
	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		if !iSValueShouldBeSearched(reflect.TypeOf(value).Field(i).Name) {
			continue
		}
		v := reflect.ValueOf(value).Field(i)
		if matched, _ := regexp.MatchString(a.reserveString[0], v.String()); matched {
			a.reserveString = append(a.reserveString, "/project/"+reflect.ValueOf(value).FieldByName("ProjectNumber").String())
			return false
		}
	}
	return true
}

func (a *copy) rangeUserForURL(key, value interface{}) bool {
	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		if !iSValueShouldBeSearched(reflect.TypeOf(value).Field(i).Name) {
			continue
		}
		v := reflect.ValueOf(value).Field(i)
		if matched, _ := regexp.MatchString(a.reserveString[0], v.String()); matched {
			a.reserveString = append(a.reserveString, "/user/"+reflect.ValueOf(value).FieldByName("Username").String())
			return false
		}
	}
	return true
}

func (a *copy) rangeClientForURL(key, value interface{}) bool {
	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		if !iSValueShouldBeSearched(reflect.TypeOf(value).Field(i).Name) {
			continue
		}
		v := reflect.ValueOf(value).Field(i)
		if matched, _ := regexp.MatchString(a.reserveString[0], v.String()); matched {
			a.reserveString = append(a.reserveString, "/client/"+reflect.ValueOf(value).FieldByName("ClientID").String())
			return false
		}
	}
	return true
}

func (a *copy) rangeDocumentForURL(key, value interface{}) bool {
	for i := 0; i < reflect.ValueOf(value).NumField(); i++ {
		if !iSValueShouldBeSearched(reflect.TypeOf(value).Field(i).Name) {
			continue
		}
		v := reflect.ValueOf(value).Field(i)
		if matched, _ := regexp.MatchString(a.reserveString[0], v.String()); matched {
			// Get relevant record in CompanyDocumentResourceCategories Table.
			c, ok := a.context.GetCompanyDocumentResourceCategoriesTable().Load(reflect.ValueOf(value).FieldByName("CategoryID").String())
			if !ok {
				return true
			}
			a.reserveString = append(a.reserveString, "/static/"+reflect.ValueOf(c).FieldByName("CategoryFolderPath").String()+"/"+reflect.ValueOf(value).FieldByName("FileName").String())
			return false
		}
	}
	return true
}

func trimUnexpectedCharactors(text string) string {
	re := regexp.MustCompile("[0-9A-Za-z\\s]")
	return strings.Join(re.FindAllString(text, -1), "")
}
