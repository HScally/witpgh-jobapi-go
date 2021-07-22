package jobmanagement

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type EmployerJobRepository struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) *EmployerJobRepository {
	return &EmployerJobRepository{db}
}

func (repository *EmployerJobRepository) AddNewEmployerJob(
	publicId string, 
	employerId int,
	EmployerCompanyId int,
	JobTitle string,
	CompanyHq string,
	JobTypeId int,
	RegionalRestrictionId int,
	JobDescription		  string
	) (*EmployerJob , error)

	var jobId int
	stmt, err := repository.db.Prepare(`INSERT into employers_jobs ("public_id", "employer_id", "employer_company_id", "is_closed", "job_title", "company_hq", "job_type_id", "regional_restriction_id", "job_description", ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`)
	
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmtErr := stmt.QueryRow(publicId, employerId, employerCompanyId, 0, JobTitle, CompanyHq, JobTypeId, RegionalRestrictionId, JobDescription).Scan(&employerId)

	if stmtErr != nil {
		return nil, err
	}

	return repository.GetEmployerJobById(employerId)
}

func (repository *EmployerJobRepository) GetEmployerJobById(employerJobId int) (*EmployerJob, error) {
	var result EmployerJob
	err := repository.db.QueryRow("select
		id,
		public_id,
		employer_id,
		employer_company_id,
		is_closed,
		job_title,
		company_hq,
		job_type_id,
		regional_restriction_id,
		job_description,
	from employers_jobs where id = $1 limit 1", employerJobId).Scan(
		&result.Id, 
		&result.PublicId, 
		&result.EmployerId, 
		&result.employerCompanyId, 
		&result.isClosed, 
		&result.JobTitle, 
		&result.CompanyHq, 
		&result.JobTypeId, 
		&result.RegionalRestrictionId, 
		&result.JobDescription)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			log.Println(err)
		}
	}

	return &result, err
}

func (repository *EmployerJobRepository) GetEmployerJobsByEmployerId(employerId int) (*EmployerJob, error) {
	employerJobs = make([]EmployerJob, 0)

	rows, err := repository.db.Prepare("select
			id,
			public_id,
			employer_id,
			employer_company_id,
			is_closed,
			job_title,
			company_hq,
			job_type_id,
			regional_restriction_id,
			job_description,
		from employers_jobs
		where employer_id = $1 limit 1", employerId).Scan(
		&result.Id, 
		&result.PublicId, 
		&result.EmployerId, 
		&result.employerCompanyId, 
		&result.isClosed, 
		&result.JobTitle, 
		&result.CompanyHq, 
		&result.JobTypeId, 
		&result.RegionalRestrictionId, 
		&result.JobDescription)
	
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(1)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var employerJob EmployerJob

	for rows.Next() {
        employerJobs = append(employerJobs, user)

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}	

