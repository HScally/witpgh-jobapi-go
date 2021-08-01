package companymanagement_test

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

func TestAddNewEmployerCompany(t *testing.T) {
	os.Clearenv()
	err := godotenv.Load("/Users/jschmitz/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	} else {
		database.ConnectWITJobBoard()

		assert := assert.New(t)

		var myRepository = repositories.NewRepositoryRegistry().GetEmployerCompanyRepository()

		var genService = generation.NewGenerationService()
		var publicId = genService.GeneratePublicId()

		p, err := myRepository.AddNewEmployerCompany(
			publicId,
			47,
			"Company, Inc.",
			"Comapny Statment",
			"",
			"companyinc.com",
			"Web Development Company",
			"123 Main St.",
			"",
			"Seattle",
			"WA",
			"98101",
			"USA",
			)

		assert.Nil(err)
		assert.NotNil(p)

		assert.Equal("Company, Inc.", p.CompanyName)
		assert.Equal("Comapny Statment", p.CompanyStatment)
		assert.Equal("companyinc.com", p.CompanyUrl)
		assert.Equal("Web Development Company", p.CompanyDescription)

		assert.Equal(0, p.HasLogo)

		assert.Equal("123 Main St.", p.AddressLine1)
		assert.Equal("Seattle", p.City)
		assert.Equal("WA", p.State)
		assert.Equal("98101", p.Zipcode)
		assert.Equal("USA", p.Country)
		fmt.Println(p.Id)
	}
}