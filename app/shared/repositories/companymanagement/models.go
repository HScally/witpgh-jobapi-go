package companymanagement

type EmployerCompanies struct {
	Id					int    `json:"id"`
	PublicId			string `json:"public_id"`
	EmployerId			int    `json:"employer_id"`
	Status				int    `json:"status"`
	CompanyName			string `json:"company_name"`
	CompanyStatment		string `json:"company_statement"`
	LogoImageFileName	string `json:"logo_image_file_name"`
	LogoObjectFileName	string `json:"logo_object_file_name"`
	CompanyLogo			string `json:"company_logo"`
	HasLogo				int    `json:"has_logo"`
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