package orm

// Note: struct name should be set as Abcabc instead of AbcAbc
// because AbcAbc will be mapped as abc_abc by gorm.

// News is News is table.
type News struct {
	ArticleID       string `gorm:"primary_key;type:int;column:ArticleID"`
	PostedBy        string `gorm:"type:nvarchar(50);column:PostedBy"`
	DateTimePosted  string `gorm:"type:datetime;column:DateTimePosted"`
	ArticleHeadline string `gorm:"type:nvarchar(MAX);column:ArticleHeadline"`
	ArticleText     string `gorm:"type:nvarchar(MAX);column:ArticleText"`
}

// Companydocumentresources is CompanyDocumentResources is table.
type Companydocumentresources struct {
	FileName         string `gorm:"primary_key;type:nvarchar(MAX);column:FileName"`
	FileFriendlyName string `gorm:"type:nvarchar(MAX);column:FileFriendlyName"`
	FileRevision     string `gorm:"type:nvarchar(10);column:FileRevision"`
	AuthorizedBy     string `gorm:"type:nvarchar(MAX);column:AuthorizedBy"`
	AuthorizedDate   string `gorm:"type:datetime;column:AuthorizedDate"`
	DownloadCount    string `gorm:"type:int;column:DownloadCount"`
	CategoryID       string `gorm:"type:int;column:CategoryID"`
}

// Companydocumentresourcecategories is CompanyDocumentResourceCategories is table.
type Companydocumentresourcecategories struct {
	CategoryID         string `gorm:"primary_key;type:int;column:CategoryID"`
	CategoryName       string `gorm:"type:nvarchar(MAX);column:CategoryName"`
	CategoryFolderPath string `gorm:"type:nvarchar(MAX);column:CategoryFolderPath"`
}

// Helpdeskticketattachedfiles is HelpDeskTicketAttachedFiles is table.
type Helpdeskticketattachedfiles struct {
	FileID   string `gorm:"primary_key;type:int;column:FileID"`
	TicketID string `gorm:"type:int;column:TicketID"`
	FileName string `gorm:"type:nvarchar(MAX);column:FileName"`
}

// Helpdesktickets is HelpDeskTickets is table.
type Helpdesktickets struct {
	TicketID string `gorm:"primary_key;type:int;column:TicketID"`

	// Status: 1 open, 0 close.
	Status string `gorm:"type:int;column:Status"`

	// Situation: -1 Low, 0 Normal, 1 Emergency.
	Situation   string `gorm:"type:int;column:Situation"`
	Username    string `gorm:"type:nvarchar(50);column:Username"`
	Computer    string `gorm:"type:nvarchar(50);column:Computer"`
	Description string `gorm:"type:nvarchar(MAX);column:Description"`
}

// Projectlocations is ProjectLocations is table.
type Projectlocations struct {
	ProjectLocationCode string `gorm:"primary_key;type:nvarchar(MAX);column:ProjectLocationCode"`
	ProjectLocationName string `gorm:"type:nvarchar(MAX);column:ProjectLocationName"`
}

// Projecttypes is ProjectTypes is table.
type Projecttypes struct {
	ProjectTypeCode string `gorm:"primary_key;type:nvarchar(MAX);column:ProjectTypeCode"`
	ProjectTypeName string `gorm:"type:nvarchar(MAX);column:ProjectTypeName"`
}

// Projectstatuss is ProjectStatuss is table.
type Projectstatuss struct {
	ProjectStatusCode string `gorm:"primary_key;type:nvarchar(MAX);column:ProjectStatusCode"`
	ProjectStatusName string `gorm:"type:nvarchar(MAX);column:ProjectStatusName"`
}

// Divisions is Divisions is table.
type Divisions struct {
	DivisionCode     string `gorm:"primary_key;type:int;column:DivisionCode"`
	DivisionName     string `gorm:"type:nvarchar(MAX);column:DivisionName"`
	FriendlyCode     string `gorm:"type:nvarchar(MAX);column:FriendlyCode"`
	IntroductoryText string `gorm:"type:nvarchar(MAX);column:IntroductoryText"`
	OfficeCode       string `gorm:"type:nvarchar(MAX);column:OfficeCode"`
	ManagerUsername  string `gorm:"type:nvarchar(MAX);column:ManagerUsername"`
}

// Offices is Offices is table.
type Offices struct {
	OfficeCode      string `gorm:"primary_key;type:nvarchar(MAX);column:OfficeCode"`
	Name            string `gorm:"type:nvarchar(MAX);column:Name"`
	StreetAddress   string `gorm:"type:nvarchar(MAX);column:StreetAddress"`
	StreetSuburb    string `gorm:"type:nvarchar(MAX);column:StreetSuburb"`
	StreetPostCode  string `gorm:"type:nvarchar(4);column:StreetPostCode"`
	PostalAddress   string `gorm:"type:nvarchar(MAX);column:PostalAddress"`
	PostalSuburb    string `gorm:"type:nvarchar(MAX);column:PostalSuburb"`
	PostalPostCode  string `gorm:"type:nvarchar(4);column:PostalPostCode"`
	State           string `gorm:"type:nvarchar(3);column:State"`
	MainPhoneNumber string `gorm:"type:nvarchar(10);column:MainPhoneNumber"`
	MainFaxNumber   string `gorm:"type:nvarchar(10);column:MainFaxNumber"`
}

// Projects is Projects table.
type Projects struct {
	ProjectNumber         string `gorm:"primary_key;type:nvarchar(MAX);column:ProjectNumber"`
	ProjectName           string `gorm:"type:nvarchar(MAX);column:ProjectName"`
	ClientID              string `gorm:"type:int;column:ClientID"`
	Division              string `gorm:"type:int;column:Division"`
	ClientRepName         string `gorm:"type:nvarchar(MAX);column:ClientRepName"`
	ClientRepTelephone    string `gorm:"type:nvarchar(MAX);column:ClientRepTelephone"`
	ClientRepMobile       string `gorm:"type:nvarchar(MAX);column:ClientRepMobile"`
	ClientRepEmailAddress string `gorm:"type:nvarchar(MAX);column:ClientRepEmailAddress"`
	ProjectDirector       string `gorm:"type:nvarchar(MAX);column:ProjectDirector"`
	ProjectManager        string `gorm:"type:nvarchar(MAX);column:ProjectManager"`
	ProjectStatusCode     string `gorm:"type:nvarchar(MAX);column:ProjectStatusCode"`
	ProjectStartDate      string `gorm:"type:datetime;column:ProjectStartDate"`
	ProjectEndDate        string `gorm:"type:datetime;column:ProjectEndDate"`
	ProjectTypeCode       string `gorm:"type:nvarchar(MAX);column:ProjectTypeCode"`
	ProjectAddress        string `gorm:"type:nvarchar(MAX);column:ProjectAddress"`
	ProjectSuburb         string `gorm:"type:nvarchar(MAX);column:ProjectSuburb"`
	ProjectLocationCode   string `gorm:"type:nvarchar(MAX);column:ProjectLocationCode"`
	ProjectDescription    string `gorm:"type:nvarchar(MAX);column:ProjectDescription"`
	ProjectValue          string `gorm:"type:decimal(18,0);column:ProjectValue"`
	ArchiveLocation       string `gorm:"type:nvarchar(MAX);column:ArchiveLocation"`
}

// Clients is Clients table.
type Clients struct {
	ClientID              string `gorm:"primary_key;type:int;column:ClientID"`
	ClientOfficeCode      string `gorm:"type:nvarchar(MAX);column:ClientOfficeCode"`
	ClientName            string `gorm:"type:nvarchar(MAX);column:ClientName"`
	ClientABNNumber       string `gorm:"type:nvarchar(MAX);column:ClientABNNumber"`
	ClientACNNumber       string `gorm:"type:nvarchar(MAX);column:ClientACNNumber"`
	ClientTypeCode        string `gorm:"type:nvarchar(MAX);column:ClientTypeCode"`
	FirstName             string `gorm:"type:nvarchar(MAX);column:FirstName"`
	LastName              string `gorm:"type:nvarchar(MAX);column:LastName"`
	ClientLocationCode    string `gorm:"type:int;column:ClientLocationCode"`
	StreetAddress         string `gorm:"type:nvarchar(MAX);column:StreetAddress"`
	StreetSuburb          string `gorm:"type:nvarchar(MAX);column:StreetSuburb"`
	StreetState           string `gorm:"type:nvarchar(MAX);column:StreetState"`
	StreetPostcode        string `gorm:"type:nvarchar(MAX);column:StreetPostcode"`
	PostalAddress         string `gorm:"type:nvarchar(MAX);column:PostalAddress"`
	PostalAddressSuburb   string `gorm:"type:nvarchar(MAX);column:PostalAddressSuburb"`
	PostalAddressState    string `gorm:"type:nvarchar(MAX);column:PostalAddressState"`
	PostalAddressPostcode string `gorm:"type:nvarchar(MAX);column:PostalAddressPostcode"`
	PhoneNumber           string `gorm:"type:nvarchar(MAX);column:PhoneNumber"`
	FaxNumber             string `gorm:"type:nvarchar(MAX);column:FaxNumber"`
	MobileNumber          string `gorm:"type:nvarchar(MAX);column:MobileNumber"`
	EMailAddress          string `gorm:"type:nvarchar(MAX);column:EMailAddress"`
}

// Users is Users table.
type Users struct {
	Username       string `gorm:"primary_key;type:nvarchar(50);column:Username"`
	FirstName      string `gorm:"type:nvarchar(MAX);column:FirstName"`
	LastName       string `gorm:"type:nvarchar(MAX);column:LastName"`
	PositionTitle  string `gorm:"type:nvarchar(MAX);column:PositionTitle"`
	DivisionCode   string `gorm:"type:int;column:DivisionCode"`
	PhoneExtension string `gorm:"type:nvarchar(MAX);column:PhoneExtension"`
	MobileNumber   string `gorm:"type:nvarchar(MAX);column:MobileNumber"`
	EmailAddress   string `gorm:"type:nvarchar(MAX);column:EmailAddress"`
	PhotoURL       string `gorm:"type:nvarchar(MAX);column:PhotoUrl"`
	Enabled        string `gorm:"type:smallint;column:Enabled"`
}

//Clienttypes is the ClientTypes Table
type Clienttypes struct {
	ClientTypeCode  string `gorm:"type:nvarchar(MAX);column:ClientTypeCode"`
	CodeDescription string `gorm:"type:nvarchar(50);column:CodeDescription"`
}

//Clientlocations is the ClientLocations Table
type Clientlocations struct {
	ClientLocationCode string `gorm:"type:int;column:ClientLocationCode"`
	ClientLocation     string `gorm:"type:nvarchar(50);column:ClientLocation"`
}

//Projectssiteinspections ... is the ProjectSiteInspections Table
type Projectssiteinspections struct {
	InspectionID       string `gorm:"type:int;column:InspectionId"`
	ProjectNumber      string `gorm:"type:nvarchar(7);column:ProjectNumber"`
	InspectedBy        string `gorm:"type:nvarchar(50);column:InspectedBy"`
	InspectionDateTime string `gorm:"type:datetime;column:InspectionDateTime"`
	InspectionDetails  string `gorm:"type:nvarchar(max);column:InspectionDetails"`
}
