package accountmanagement


type EmployerCompanies struct {
	Id					int    `json:"id"`	
	PublicId			string `json:"public_id"`
	EmployerId			int    `json:"employer_id"`
	Status				bool   `json:"status"`
	CompanyName			string `json:"company_name"`
	CompanyStatment		string `json:"company_statement"`
	LogoImageFileName	string `json:"logo_image_file_name"`
	LogoObjectFileName	string `json:"logo_object_file_name"`
	CompanyLogo			string `json:"company_logo"`	
	HasLogo				bool   `json:"has_logo"`
	CompanyUrl			string `json:"company_url"`
	CompanyDescription	string `json:"company_description"`
	AddressLine1		string `json:"address_line1"`
	AddressLine2		string `json:"address_line2"`
	City				string `json:"city"`
	State				string `json:"state"`
	Zipcode				string `json:"zipcode"`
	Country				string `json:"country"`
	Latitude			int    `json:latitude`
	Longitude			int	   `json:"longitude"`
}

type Employer struct {
	Id                int    `json:"id"`
	PublicId          string `json:"public_id"`
	EmployerKey       string `json:"employer_key"`
	Status            int    `json:"status"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	MustResetPassword int    `json:"must_reset_password"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	SendMobileNotices int    `json:"send_mobile_notices"`
}

type JobRegionalRestrictions struct {
	Id			  int 	 `json:"id"`
	Description   string `json:"description"`
	RestrictionId int    `json:"restriction_id"`
}

type JobTypes struct {
	Id			int    `json:"id"`
	Description string `json:"description"`
	TypeId		int    `json:"type_id"`
}