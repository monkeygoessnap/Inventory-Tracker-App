/*
Package env provides the initialization of the env file to be used as variables
*/
package env

import (
	"PGL/Client/log"

	"github.com/joho/godotenv"
)

//initializes the env file
func InitEnv() {
	if err := godotenv.Load("./env/.env"); err != nil {
		log.Error.Fatal(err)
	} else {
		log.Info.Println("Successfully loaded .env file")
	}
}
