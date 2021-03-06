package accountmanagement

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (repository *AccountRepository) AddNewEmployer(publicId string, employerKey string, email string, password string, firstname string, lastname string) (*Employer, error) {
	var employerId int
	stmt, err := repository.db.Prepare(`INSERT into employers (public_id, employer_key, status, email, password, must_reset_password, firstname, lastname, send_mobile_notices) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmtErr := stmt.QueryRow(publicId, employerKey, 1, email, password, 1, firstname, lastname, 0).Scan(&employerId)

	if stmtErr != nil {
		log.Println(stmtErr)
		return nil, err
	}

	return repository.GetEmployerById(employerId)
}

func (repository *AccountRepository) GetEmployerById(employerId int) (*Employer, error) {
	var result Employer

	err := repository.db.QueryRow(`select id, public_id, employer_key, status, email, password, must_reset_password, firstname, lastname, send_mobile_notices from employers where id = $1 limit 1`, employerId).Scan(&result.Id, &result.PublicId, &result.EmployerKey, &result.Status, &result.Email, &result.Password, &result.MustResetPassword, &result.Firstname, &result.Lastname, &result.SendMobileNotices)

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

func (repository *AccountRepository) UpdateEmployerPassword(employerKey string, password string) (*Employer, error) {
	var employerId int
	stmt, err := repository.db.Prepare(`UPDATE employers SET password = $2 WHERE employer_key = $1 RETURNING id`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmtErr := stmt.QueryRow(employerKey, password).Scan(&employerId)

	if stmtErr != nil {
		log.Println(stmtErr)
		return nil, err
	}

	return repository.GetEmployerById(employerId)
}

func (repository *AccountRepository) UpdateEmployerEmail(employerKey string, email string) (*Employer, error) {
	var employerId int
	stmt, err := repository.db.Prepare(`UPDATE employers SET email = $2 WHERE employer_key = $1 RETURNING id`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmtErr := stmt.QueryRow(employerKey, email).Scan(&employerId)

	if stmtErr != nil {
		log.Println(stmtErr)
		return nil, err
	}

	return repository.GetEmployerById(employerId)
}

func (repository *AccountRepository) UpdateEmployerMobileNotice(employerKey string, sendMobileNotices int) (*Employer, error) {
	var employerId int
	// TODO: Refactor all these functions to remove duplciation/DRY
	// We can also write a private function to construct the SQL command
	stmt, err := repository.db.Prepare(`UPDATE employers SET send_mobile_notices = $2 WHERE employer_key = $1 RETURNING id`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	stmtErr := stmt.QueryRow(employerKey, sendMobileNotices).Scan(&employerId)

	if stmtErr != nil {
		log.Println(stmtErr)
		return nil, err
	}

	return repository.GetEmployerById(employerId)
}