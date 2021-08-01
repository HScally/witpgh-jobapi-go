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