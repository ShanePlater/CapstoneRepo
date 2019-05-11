package types

const (
	// SuggestionSubject represents subject of suggestion request email.
	SuggestionSubject = "New suggestion request from "

	// SuggestionTmpl is html template of suggestion request email.
	SuggestionTmpl = `
	{{ define "email" }}
	<p>{{ .Text  }}</p>
	<br>
	<p>Created by {{ .Author }}</p>
	{{ end }}
	`

	// TicketSubject represents subject of help desk ticket email.
	TicketSubject = "New Help Desk Ticket #"

	// TicketTmpl is html template of help desk ticket email.
	TicketTmpl = `
	{{ define "email" }}
	<p>{{ .Text  }}</p>
	<br>
	<p>Created by {{ .Author }}</p>
	{{ end }}
	`
)

// News represents a news entity shown on home page.
type News struct {
	ID     int
	Author string
	Date   string
	Title  string
	Text   string
}

// GetNewsJSON is the body of getNews API request.
type GetNewsJSON struct {
	StartDate string
	EndDate   string
	Length    int
}

// Resource represents a document resource entity.
type Resource struct {
	FileID         string
	FileName       string
	FileRevision   string
	AuthorizedBy   string
	AuthorizedDate string
	CategoryID     int
	URL            string
}

// RecentJSON is the body of recent API request.
type RecentJSON struct {
	Length int
}

// GetResourcesJSON is the body of getResources API request.
type GetResourcesJSON struct {
	CategoryID int
}

// Option represents any data in form of ID and Name.
type Option struct {
	ID   string
	Name string
}

// File represents an attached file for help desk ticket.
type File struct {
	FileID   int
	FileName string
	URL      string
}

// CreateSuggestionJSON is the body of createSuggestion API request.
type CreateSuggestionJSON struct {
	Author string
	Text   string
}

// HelpDeskTicket represents a help desk ticket.
type HelpDeskTicket struct {
	TicketID    int
	Situation   int
	Status      int
	Username    string
	Computer    string
	Description string
	Files       []File
}

// HelpDeskTicketInterface represents a help desk ticket.
// It is designed for utilizing by models.
type HelpDeskTicketInterface struct {
	TicketID    int
	Situation   int
	Status      int
	Username    string
	Computer    string
	Description string
	Files       []interface{}
}

// SearchRedbookJSON is the body of searchRedbook API request.
type SearchRedbookJSON struct {
	ClientName string
	Address    string
	Suburb     string
	Location   string
	Type       string
	Status     string
	Division   string
	Office     string
	StartDate  string
	EndDate    string
}

// Project represent a project entity in search results.
type Project struct {
	ID         string
	Name       string
	ClientName string
	Company    string
	Address    string
	Division   string
}
type Project2 struct {
	ID         	string
	Name       	string
	ClientName 	string	
	Address    	string
	Suburb	   	string
	Location   	string
	Type       	string
	Status     	string
	StartDate  	string
	EndDate    	string
	//CR = Client Rep
	CRName     	string
	CRPhone    	string
	CRMobile   	string
	CREmail    	string
	Division   	string
	Director   	string
	Manager    	string
	Value       string
	Description string

}

// SearchJSON is the body of searchProjects, searchUsers,
// searchClients, and searchResources API request.
type SearchJSON struct {
	Keyword string
}

// User represent a user entity in search results.
type User struct {
	ID             string
	Name           string
	PhoneNumber    string
	PhoneExtension string
	Email          string
	Division       int
}

// Client represent a client entity in search results.
type Client struct {
	ID          string
	Name        string
	PhoneNumber string
	FaxNumber   string
	Address     string
}

type GetByIDJSON struct {
	ID string
}

type GetDivisionsJSON struct {
	Office   string
	Division string
}

type Division struct {
	EmailAddress   string
	Name           string
	MobileNumber   string
	PhoneExtension string
	PhotoURL       string
	PositionTitle  string
	Username       string
}

type GetMeJSON struct {
	Username string
	Password string
}
