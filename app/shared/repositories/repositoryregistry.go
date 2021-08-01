
package repositories

import (
	"witpgh-jobapi-go/app/shared/database"
	"witpgh-jobapi-go/app/shared/repositories/accountmanagement"
	"witpgh-jobapi-go/app/shared/repositories/jobmanagement"
	"witpgh-jobapi-go/app/shared/repositories/companymanagement"
)

type RepositoryRegistry struct {
}

func NewRepositoryRegistry() *RepositoryRegistry {
	return &RepositoryRegistry{}
}

func (registry *RepositoryRegistry) GetEmployerAccountRepository() *accountmanagement.AccountRepository {
	return accountmanagement.NewAccountRepository(database.WITPGH)
}

func (registry *RepositoryRegistry) GetEmployerJobRepository() *jobmanagement.EmployerJobRepository {
	return jobmanagement.NewJobRepository(database.WITPGH)
}

func (registry *RepositoryRegistry) GetEmployerCompanyRepository() *companymanagement.EmployerCompanyRepository {
	return companymanagement.NewEmployerCompanyRepository(database.WITPGH)
}