package companymanagement

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type EmployerCompanyRepository struct {
	db *sql.DB
}

func NewEmployerCompanyRepository(db *sql.DB) *EmployerCompanyRepository {
	return &EmployerCompanyRepository{db}
}

func (repository *EmployerCompanyRepository) AddNewEmployerCompany(
	publicId string,
	employerId string,
	companyName string,
	companyStatment	string,
	companyLogo string,
	companyUrl string,
	companyDescription string,
	addressLine1 string,
	addressLine2 string,
	city string,
	state string,
	zipcode string,
	country string) (*EmployerCompanies, error) {
	var employerCompanyId int
	stmt, err := repository.db.Prepare(
		`INSERT into employer_companies
		(public_id, employer_id, status, company_name, company_statement, company_logo, has_logo, company_url, company_description, address_line1, address_line2, city, state, zipcode, country, latitude, longitude)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17)
		RETURNING id`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var hasLogo int
	if (companyLogo != nil && companyLogo != '') {
		hasLogo = 1
	} else {
		hasLogo = 0
	}

	stmtErr := stmt.QueryRow(publicId,
		employerId,
		companyName,
		1,
		companyStatment,
		companyLogo,
		hasLogo,
		companyUrl,
		companyDescription,
		addressLine1,
		addressLine2,
		city,
		state,
		zipcode,
		country,
		0.0,
		0.0).Scan(&employerCompanyId)

	if stmtErr != nil {
		log.Println(stmtErr)
		return nil, err
	}

	return repository.GetEmployerCompanyById(employerCompanyId)
}

func (repository *EmployerCompanyRepository) GetEmployerCompanyById(employerCompanyId int) (*EmployerCompanies, error) {
	var result EmployerCompanies

	err := repository.db.QueryRow(`
	select id,
		public_id,
		employer_id,
		status,
		company_name,
		company_statement,
		company_logo,
		has_logo,
		company_url,
		company_description,
		address_line1,
		address_line2,
		city,
		state,
		zipcode,
		country,
		latitude,
		longitude
		where id = $1 limit 1`, employerCompanyId).Scan(
			&result.Id,
			&result.PublicId,
			&result.EmployerId,
			&result.Status,
			&result.CompanyName,
			&result.CompanyStatment,
			&result.CompanyLogo,
			&result.HasLogo,
			&result.CompanyUrl,
			&result.CompanyDescription,
			&result.AddressLine1,
			&result.AddressLine2,
			&result.City,
			&result.State,
			&result.Zipcode,
			&result.Country,
			&result.Latitude,
			&result.Longitude)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			return nil, err
		} else {
			log.Println(err)
		}
	}

	return &result, err
}