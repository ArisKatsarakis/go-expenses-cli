package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Errror: ", err.Error())
	}

	return os.Getenv(key)

}
