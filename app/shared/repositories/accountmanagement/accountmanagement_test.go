package accountmanagement_test

import (
	"witpgh-jobapi-go/app/shared/database"
	"witpgh-jobapi-go/app/shared/repositories"
	"witpgh-jobapi-go/app/shared/services/system/generation"

	"os"
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAddNewEmployer(t *testing.T) {
	os.Clearenv()
	err := godotenv.Load("/Users/jschmitz/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	} else {
		database.ConnectWITJobBoard()

		assert := assert.New(t)

		var myRepository = repositories.NewRepositoryRegistry().GetEmployerAccountRepository()

		// Generate random digits for each primary key
		// TODO: Move this to AddNewEmployer in `accountmanagement.go`
		var genService = generation.NewGenerationService()
		var publicId = genService.GeneratePublicId()
		var companyId = genService.GeneratePublicId()
		var emailAddressBody = genService.GeneratePublicId()
		email := fmt.Sprintf("%s@test.com", emailAddressBody)

		p, err := myRepository.AddNewEmployer(publicId, companyId, email, "pasword123", "Test", "Client")

		assert.Nil(err)
		assert.NotNil(p)
		assert.True(p.Id > 1)
		assert.True(p.PublicId == publicId)
	}
}

func TestGetEmployerById(t *testing.T) {
	os.Clearenv()
	err := godotenv.Load("/Users/jschmitz/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	} else {
		database.ConnectWITJobBoard()

		assert := assert.New(t)

		var myRepository = repositories.NewRepositoryRegistry().GetEmployerAccountRepository()

		p, err := myRepository.GetEmployerById(47)

		assert.Nil(err)
		assert.NotNil(p)
		assert.NotNil(p.Email)
		assert.Equal("pasword123", p.Password)
		assert.Equal("Test", p.Firstname)
		assert.Equal("Client", p.Lastname)
	}
}

func TestUpdateEmployerPassword(t *testing.T) {
	os.Clearenv()
	err := godotenv.Load("/Users/jschmitz/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	} else {
		database.ConnectWITJobBoard()

		assert := assert.New(t)

		var myRepository = repositories.NewRepositoryRegistry().GetEmployerAccountRepository()

		p, err := myRepository.UpdateEmployerPassword("d7d1a21d7cccb92c71d5bfdcd20c77fc", "newpassword123")

		assert.Nil(err)
		assert.NotNil(p)
		assert.Equal("newpassword123", p.Password)
	}
}

func TestUpdateEmployerEmail(t *testing.T) {
	os.Clearenv()
	err := godotenv.Load("/Users/jschmitz/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	} else {
		database.ConnectWITJobBoard()

		assert := assert.New(t)

		var myRepository = repositories.NewRepositoryRegistry().GetEmployerAccountRepository()

		p, err := myRepository.UpdateEmployerEmail("d7d1a21d7cccb92c71d5bfdcd20c77fc", "NewTestEmail@test.com")

		assert.Nil(err)
		assert.NotNil(p)
		assert.Equal("NewTestEmail@test.com", p.Email)
	}
}

func TestUpdateEmployerMobileNotice(t *testing.T) {
	os.Clearenv()
	err := godotenv.Load("/Users/jschmitz/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	} else {
		database.ConnectWITJobBoard()

		assert := assert.New(t)

		var myRepository = repositories.NewRepositoryRegistry().GetEmployerAccountRepository()

		p, err := myRepository.UpdateEmployerMobileNotice("d7d1a21d7cccb92c71d5bfdcd20c77fc", 0)

		assert.Nil(err)
		assert.NotNil(p)
		assert.Equal(0, p.SendMobileNotices)
	}
}