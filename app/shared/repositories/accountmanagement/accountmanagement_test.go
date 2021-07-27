package accountmanagement_test

import (
	"witpgh-jobapi-go/app/shared/database"
	"witpgh-jobapi-go/app/shared/repositories"

	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAddNewEmployer(t *testing.T) {
	os.Clearenv()
	err := godotenv.Load("/Users/jschmitz/projects/go/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	} else {
		database.ConnectWITJobBoard()

		assert := assert.New(t)

		var myRepository = repositories.NewRepositoryRegistry().GetEmployerAccountRepository()
		p, err := myRepository.AddNewEmployer("abcd", "xyz", "test@test.com", "pasword123", "Test", "Client")

		assert.Nil(err)
		assert.Equal(p, 1, "Are Equal")
		assert.True(p.PublicId == "abcd")
	}
}