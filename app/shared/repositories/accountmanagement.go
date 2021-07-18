package accountmanagement

type Employer struct {
	Id int
	PublicId string
	EmployerKey string
	Status int
	Email string
	Password string
	MustResetPassword int
	Firstname string
	Lastname string
	SendMobileNotices int
}
