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
	employerCompanyId int,
	jobTitle string,
	companyHq string,
	jobTypeId int,
	regionalRestrictionId int,
	jobDescription string) (*EmployerJob , error) {

	var jobId int
	stmt, err := repository.db.Prepare(`INSERT into employer_jobs (public_id, employer_id, employer_company_id, is_closed, job_title, company_hq, job_type_id, regional_restriction_id, job_description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmtErr := stmt.QueryRow(publicId, employerId, employerCompanyId, 0, jobTitle, companyHq, jobTypeId, regionalRestrictionId, jobDescription).Scan(&jobId)

	if stmtErr != nil {
		log.Println(stmtErr)
		return nil, err
	}

	return repository.GetEmployerJobById(jobId)
}


func (repository *EmployerJobRepository) GetEmployerJobById(employerJobId int) (*EmployerJob, error) {
	var result EmployerJob
	err := repository.db.QueryRow(`select
		id,
		public_id,
		employer_id,
		employer_company_id,
		is_closed,
		job_title,
		company_hq,
		job_type_id,
		regional_restriction_id,
		job_description
	from employer_jobs where id = $1 limit 1`, employerJobId).Scan(
		&result.Id,
		&result.PublicId,
		&result.EmployerId,
		&result.EmployerCompanyId,
		&result.IsClosed,
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

// func (repository *EmployerJobRepository) GetEmployerJobsByEmployerId(employerId int) (*EmployerJob, error) {
// 	employerJobs = make([]EmployerJob, 0)

// 	rows, err := repository.db.Prepare(`select
// 			id,
// 			public_id,
// 			employer_id,
// 			employer_company_id,
// 			is_closed,
// 			job_title,
// 			company_hq,
// 			job_type_id,
// 			regional_restriction_id,
// 			job_description,
// 		from employers_jobs
// 		where employer_id = $1 limit 1`, employerId).Scan(
// 		&result.Id,
// 		&result.PublicId,
// 		&result.EmployerId,
// 		&result.employerCompanyId,
// 		&result.isClosed,
// 		&result.JobTitle,
// 		&result.CompanyHq,
// 		&result.JobTypeId,
// 		&result.RegionalRestrictionId,
// 		&result.JobDescription)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer stmt.Close()

// 	rows, err = stmt.Query(1)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()

// 	var employerJob EmployerJob

// 	for rows.Next() {
//         employerJobs = append(employerJobs, user)

// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	}
// }

// func (repository *EmployerJobRepository) GetEmployerRemoteJobs() (*EmployerJob, error) {
// 	employerJobs = make([]EmployerJob, 0)

// 	rows, err := repository.db.Prepare(`select
// 			id,
// 			public_id,
// 			employer_id,
// 			employer_company_id,
// 			is_closed,
// 			job_title,
// 			company_hq,
// 			job_type_id,
// 			regional_restriction_id,
// 			job_description,
// 		from employers_jobs
// 		where regional_restriction_id = 1`).Scan(
// 		&result.Id,
// 		&result.PublicId,
// 		&result.EmployerId,
// 		&result.employerCompanyId,
// 		&result.isClosed,
// 		&result.JobTitle,
// 		&result.CompanyHq,
// 		&result.JobTypeId,
// 		&result.RegionalRestrictionId,
// 		&result.JobDescription)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer stmt.Close()

// 	rows, err = stmt.Query(1)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()

// 	var employerJob EmployerJob

// 	for rows.Next() {
//         employerJobs = append(employerJobs, user)

// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	}
// }

// func (repository *EmployerJobRepository) GetEmployerJobsByRegion(regionalRestrictionId int) (*EmployerJob, error) {
// 	employerJobs = make([]EmployerJob, 0)

// 	rows, err := repository.db.Prepare(`select
// 			id,
// 			public_id,
// 			employer_id,
// 			employer_company_id,
// 			is_closed,
// 			job_title,
// 			company_hq,
// 			job_type_id,
// 			regional_restriction_id,
// 			job_description,
// 		from employers_jobs
// 		where regional_restriction_id = $1`, regionalRestrictionId).Scan(
// 		&result.Id,
// 		&result.PublicId,
// 		&result.EmployerId,
// 		&result.employerCompanyId,
// 		&result.isClosed,
// 		&result.JobTitle,
// 		&result.CompanyHq,
// 		&result.JobTypeId,
// 		&result.RegionalRestrictionId,
// 		&result.JobDescription)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer stmt.Close()

// 	rows, err = stmt.Query(1)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()

// 	var employerJob EmployerJob

// 	for rows.Next() {
//         employerJobs = append(employerJobs, user)

// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	}
// }

// func (repository *EmployerJobRepository) GetEmployerOpenJobsByEmployerId(employerId int) (*EmployerJob, error) {
// 	employerJobs = make([]EmployerJob, 0)

// 	rows, err := repository.db.Prepare(`select
// 			id,
// 			public_id,
// 			employer_id,
// 			employer_company_id,
// 			is_closed,
// 			job_title,
// 			company_hq,
// 			job_type_id,
// 			regional_restriction_id,
// 			job_description,
// 		from employers_jobs
// 		where is_closed = 0 and employer_id = $1 limit 1`, employerId).Scan(
// 		&result.Id,
// 		&result.PublicId,
// 		&result.EmployerId,
// 		&result.employerCompanyId,
// 		&result.isClosed,
// 		&result.JobTitle,
// 		&result.CompanyHq,
// 		&result.JobTypeId,
// 		&result.RegionalRestrictionId,
// 		&result.JobDescription)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer stmt.Close()

// 	rows, err = stmt.Query(1)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()

// 	var employerJob EmployerJob

// 	for rows.Next() {
//         employerJobs = append(employerJobs, user)

// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	}
// }

// func (repository *EmployerJobRepository) GetEmployerOpenRemoteJobs() (*EmployerJob, error) {
// 	employerJobs = make([]EmployerJob, 0)

// 	rows, err := repository.db.Prepare(`select
// 			id,
// 			public_id,
// 			employer_id,
// 			employer_company_id,
// 			is_closed,
// 			job_title,
// 			company_hq,
// 			job_type_id,
// 			regional_restriction_id,
// 			job_description,
// 		from employers_jobs
// 		where is_closed = 0 and regional_restriction_id = 1`).Scan(
// 		&result.Id,
// 		&result.PublicId,
// 		&result.EmployerId,
// 		&result.employerCompanyId,
// 		&result.isClosed,
// 		&result.JobTitle,
// 		&result.CompanyHq,
// 		&result.JobTypeId,
// 		&result.RegionalRestrictionId,
// 		&result.JobDescription)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer stmt.Close()

// 	rows, err = stmt.Query(1)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()

// 	var employerJob EmployerJob

// 	for rows.Next() {
//         employerJobs = append(employerJobs, user)

// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	}
// }

// func (repository *EmployerJobRepository) GetEmployerOpenJobsByRegion(regionalRestrictionId int) (*EmployerJob, error) {
// 	employerJobs = make([]EmployerJob, 0)

// 	rows, err := repository.db.Prepare(`select
// 			id,
// 			public_id,
// 			employer_id,
// 			employer_company_id,
// 			is_closed,
// 			job_title,
// 			company_hq,
// 			job_type_id,
// 			regional_restriction_id,
// 			job_description,
// 		from employers_jobs
// 		where is_closed = 0 and regional_restriction_id = $1`, regionalRestrictionId).Scan(
// 		&result.Id,
// 		&result.PublicId,
// 		&result.EmployerId,
// 		&result.employerCompanyId,
// 		&result.isClosed,
// 		&result.JobTitle,
// 		&result.CompanyHq,
// 		&result.JobTypeId,
// 		&result.RegionalRestrictionId,
// 		&result.JobDescription)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer stmt.Close()

// 	rows, err = stmt.Query(1)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer rows.Close()

// 	var employerJob EmployerJob

// 	for rows.Next() {
//         employerJobs = append(employerJobs, user)

// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	}
// }

// func (repository *EmployerJobRepository) CloseEmployerJob(employerJobId int) (*EmployerJob, error) {
// 	var result EmployerJob

// 	stmt, err := repository.db.Prepare(`Update employers_jobs ("public_id", "employer_id", "employer_company_id", "is_closed", "job_title", "company_hq", "job_type_id", "regional_restriction_id", "job_description", ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`)

// 	// err := repository.db.QueryRow(`select
// 	// 	id,
// 	// 	public_id,
// 	// 	employer_id,
// 	// 	employer_company_id,
// 	// 	is_closed,
// 	// 	job_title,
// 	// 	company_hq,
// 	// 	job_type_id,
// 	// 	regional_restriction_id,
// 	// 	job_description,
// 	// from employers_jobs where id = $1 limit 1`, employerJobId).Scan(
// 	// 	&result.Id,
// 	// 	&result.PublicId,
// 	// 	&result.EmployerId,
// 	// 	&result.employerCompanyId,
// 	// 	&result.isClosed,
// 	// 	&result.JobTitle,
// 	// 	&result.CompanyHq,
// 	// 	&result.JobTypeId,
// 	// 	&result.RegionalRestrictionId,
// 	// 	&result.JobDescription)

// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return nil, err
// 		} else {
// 			log.Println(err)
// 		}
// 	}

// 	return &result, err
// }