package mcustomer

type CustomerPost struct {
	ImpUuid       string `json:"imp_uuid"`
	ImpNome       string `json:"imp_nome"`
	ImpStatus     int    `json:"imp_status"`
	ImpId         string `json:"imp_id"`
	ImpIp         string `json:"imp_ip"`
	ImpUserCreate string `json:"imp_user_create"`
	ImpUserUpdate string `json:"imp_user_update"`

	Channel string `json:"channel"`
	UserID  string `json:"userId"`

	Customer Customer
}

type Customer struct {
	Name              string `json:"name"`
	BirthDate         string `json:"birthDate"`
	Sex               string `json:"sex"`
	MotherName        string `json:"motherName"`
	Type              string `json:"type"`
	MaritalStatus     string `json:"maritalStatus"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Phone2            string `json:"phone2"`
	Phone3            string `json:"phone3"`
	Emancipated       string `json:"emancipated"`
	IncomeReported    string `json:"incomeReported"`
	ValueInConcurrent string `json:"valueInConcurrent"`
	Nationality       string `json:"nationality"`
	RuralObligation   string `json:"ruralObligation"`
	Title             string `json:"title"`
	Illiterate        string `json:"illiterate"`
	Document          Document
	Disabilities      []Disabilities
	SocialMedias      []SocialMedias
	Address           Address
	CompanyInfo       CompanyInfo

	Contact Contact
}

type Document struct {
	Type         string `json:"type"`
	Number       string `json:"number"`
	ValidityDate string `json:"validityDate"`
	IssueDate    string `json:"issueDate"`
	IssuerAgency string `json:"issuerAgency"`
	IssuerState  string `json:"issuerState"`
}

type Disabilities struct {
	HasDisability  bool   `json:"hasDisability"`
	DisabilityKind string `json:"disabilityKind"`
}

type SocialMedias struct {
	Name  string `json:"name"`
	Login string `json:"login"`
}

type Address struct {
	StreetType      string `json:"streetType"`
	StreetName      string `json:"streetName"`
	Number          string `json:"number"`
	Complement      string `json:"complement"`
	Neighborhood    string `json:"neighborhood"`
	Zip             string `json:"zip"`
	City            string `json:"city"`
	State           string `json:"state"`
	Special         string `json:"special"`
	PostalCode      string `json:"postalCode"`
	Country         string `json:"country"`
	Reference       bool   `json:"reference"`
	Apartment       string `json:"apartment"`
	ComplementType1 string `json:"complementType1"`
	ComplementDesc1 string `json:"complementDesc1"`
	ComplementType2 string `json:"complementType2"`
	ComplementDesc2 string `json:"complementDesc2"`
	ComplementType3 string `json:"complementType3"`
	ComplementDesc3 string `json:"complementDesc3"`
}

type CompanyInfo struct {
	TaxExempted                string `json:"taxExempted"`
	StateRegistration          string `json:"stateRegistration"`
	StateRegistrationIssuer    string `json:"stateRegistrationIssuer"`
	StateRegistrationState     string `json:"stateRegistrationState"`
	StateRegistrationCity      string `json:"stateRegistrationCity"`
	MunicipalRegistration      string `json:"municipalRegistration"`
	MunicipalRegistrationCity  string `json:"municipalRegistrationCity"`
	MunicipalRegistrationState string `json:"municipalRegistrationState"`
	SocialCapital              string `json:"socialCapital"`
	Phone                      string `json:"phone"`
	Phone2                     string `json:"phone2"`
	Phone3                     string `json:"phone3"`
	Email                      string `json:"email"`
	FoundationDate             string `json:"foundationDate"`
	EmployeesNumber            string `json:"employeesNumber"`
	LegalNature                string `json:"legalNature"`
	CompanySize                string `json:"companySize"`
	EnterpriseGroup            string `json:"enterpriseGroup"`
	ActivityBranch             string `json:"activityBranch"`
	Cnae                       string `json:"cnae"`
	CnaeDesc                   string `json:"cnaeDesc"`
	CnaeSec                    string `json:"cnaeSec"`
	CnaeSecDesc                string `json:"cnaeSecDesc"`
	LegalNatureID              string `json:"legalNatureId"`
	CoverageIntelig            string `json:"coverageIntelig"`
	PresumedBilling            string `json:"presumedBilling"`
	MarketSegment              string `json:"marketSegment"`
	WalletBelonging            string `json:"walletBelonging"`
	RuralObligation            string `json:"ruralObligation"`
}

type Contact struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	BirthDate       string `json:"birthDate"`
	SocialSecNo     string `json:"socialSecNo"`
	Sex             string `json:"sex"`
	Email           string `json:"email"`
	ContactDocument ContactDocument
	ContactAddress  ContactAddress
}

type ContactDocument struct {
	Type         string `json:"type"`
	Number       string `json:"number"`
	IssuerState  string `json:"issuerState"`
	IssuerAgency string `json:"issuerAgency"`
}

type ContactAddress struct {
	StreetType   string `json:"streetType"`
	StreetName   string `json:"streetName"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	Zip          string `json:"zip"`
	City         string `json:"city"`
	State        string `json:"state"`
}
