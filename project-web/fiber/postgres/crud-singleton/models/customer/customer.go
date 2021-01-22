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
	UserID  string `json:"user_id"`

	Customer Customer
}

type Customer struct {
	Name              string `json:"name"`
	BirthDate         string `json:"birth_date"`
	Sex               string `json:"sex"`
	MotherName        string `json:"mother_name"`
	Type              string `json:"type"`
	MaritalStatus     string `json:"marital_status"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	Phone2            string `json:"phone2"`
	Phone3            string `json:"phone3"`
	Emancipated       string `json:"emancipated"`
	IncomeReported    string `json:"income_reported"`
	ValueInConcurrent string `json:"value_in_concurrent"`
	Nationality       string `json:"nationality"`
	RuralObligation   string `json:"rural_obligation"`
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
	ValidityDate string `json:"validity_date"`
	IssueDate    string `json:"issue_date"`
	IssuerAgency string `json:"issuer_agency"`
	IssuerState  string `json:"issuer_state"`
}

type Disabilities struct {
	HasDisability  bool   `json:"has_disability"`
	DisabilityKind string `json:"disability_kind"`
}

type SocialMedias struct {
	Name  string `json:"name"`
	Login string `json:"login"`
}

type Address struct {
	StreetType      string `json:"street_type"`
	StreetName      string `json:"street_name"`
	Number          string `json:"number"`
	Complement      string `json:"complement"`
	Neighborhood    string `json:"neighborhood"`
	Zip             string `json:"zip"`
	City            string `json:"city"`
	State           string `json:"state"`
	Special         string `json:"special"`
	PostalCode      string `json:"postal_code"`
	Country         string `json:"country"`
	Reference       bool   `json:"reference"`
	Apartment       string `json:"apartment"`
	ComplementType1 string `json:"complement_type1"`
	ComplementDesc1 string `json:"complement_desc1"`
	ComplementType2 string `json:"complement_type2"`
	ComplementDesc2 string `json:"complement_desc2"`
	ComplementType3 string `json:"complement_type3"`
	ComplementDesc3 string `json:"complement_desc3"`
}

type CompanyInfo struct {
	TaxExempted                string `json:"tax_exempted"`
	StateRegistration          string `json:"state_registration"`
	StateRegistrationIssuer    string `json:"state_registrationIssuer"`
	StateRegistrationState     string `json:"state_registrationState"`
	StateRegistrationCity      string `json:"state_registrationCity"`
	MunicipalRegistration      string `json:"municipal_registration"`
	MunicipalRegistrationCity  string `json:"municipal_registration_city"`
	MunicipalRegistrationState string `json:"municipal_registrationState"`
	SocialCapital              string `json:"social_capital"`
	Phone                      string `json:"phone"`
	Phone2                     string `json:"phone2"`
	Phone3                     string `json:"phone3"`
	Email                      string `json:"email"`
	FoundationDate             string `json:"foundation_date"`
	EmployeesNumber            string `json:"employees_number"`
	LegalNature                string `json:"legal_nature"`
	CompanySize                string `json:"company_size"`
	EnterpriseGroup            string `json:"enterprise_group"`
	ActivityBranch             string `json:"activity_branch"`
	Cnae                       string `json:"cnae"`
	CnaeDesc                   string `json:"cnae_desc"`
	CnaeSec                    string `json:"cnae_sec"`
	CnaeSecDesc                string `json:"cnae_sec_desc"`
	LegalNatureID              string `json:"legal_natureId"`
	CoverageIntelig            string `json:"coverage_intelig"`
	PresumedBilling            string `json:"presumed_billing"`
	MarketSegment              string `json:"market_segment"`
	WalletBelonging            string `json:"wallet_belonging"`
	RuralObligation            string `json:"rural_obligation"`
}

type Contact struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	BirthDate       string `json:"birth_date"`
	SocialSecNo     string `json:"social_sec_no"`
	Sex             string `json:"sex"`
	Email           string `json:"email"`
	ContactDocument ContactDocument
	ContactAddress  ContactAddress
}

type ContactDocument struct {
	Type         string `json:"type"`
	Number       string `json:"number"`
	IssuerState  string `json:"issuer_state"`
	IssuerAgency string `json:"issuer_agency"`
}

type ContactAddress struct {
	StreetType   string `json:"street_type"`
	StreetName   string `json:"street_name"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	Zip          string `json:"zip"`
	City         string `json:"city"`
	State        string `json:"state"`
}
