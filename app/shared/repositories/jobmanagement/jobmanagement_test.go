package jobmanagement_test

import (
	"witpgh-jobapi-go/app/shared/database"
	"witpgh-jobapi-go/app/shared/repositories"

	"os"
	"log"
	"testing"
	"math/rand"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)
func Random(min int, max int) int {
	rand.Seed(time.Now().Unix())
	 return rand.Intn(max-min) + min
}

func TestAddNewEmployerJob(t *testing.T) {

	os.Clearenv()
	err := godotenv.Load("/Users/jschmitz/src/witpgh-jobapi-go/doc.env")

	if err != nil {
		log.Print(err)
	} else {
		database.ConnectWITJobBoard()

		assert := assert.New(t)

		var publicId = strconv.Itoa(Random(0, 1000))
		var employerId = strconv.Itoa(Random(0, 1000))

		var myRepository = repositories.NewRepositoryRegistry().GetEmployerJobRepository()
		p, err := myRepository.AddNewEmployerJob(publicId, "abcd", 1, "Programmer", "Seattle", 1, 1, "Entry Level Position")

		assert.Nil(err)
		assert.NotNil(p)
		assert.True(p.Id > 1)
		assert.True(p.PublicId == publicId)
	}

}