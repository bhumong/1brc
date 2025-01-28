package filepath

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetPath() string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("BRC1_FULL_PATH")
}
