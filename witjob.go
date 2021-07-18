package main
// go mod init
// go mod tidy
// go mod vendor
import (
	"log"
	"net/http"
	"os"
	"runtime"
	""

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Lshortfile)
	
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	
	database.ConnectWITJobBoard()

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "3000"
	}	
	http.ListenAndServe(":" + port, nil)
}

func configureEnviornemnts() {
	os.Clearenv()

	err := godotenv.Load("doc.env")
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
}
}
