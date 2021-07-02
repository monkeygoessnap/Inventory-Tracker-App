/*
Env package provides the initialization of the env file to usable variables
*/
package env

import (
	"PGL/APIServer/log"

	"github.com/joho/godotenv"
)

//Initializes the env file
func InitEnv() {
	if err := godotenv.Load("./env/.env"); err != nil {
		log.Error.Fatal(err)
	} else {
		log.Info.Println("Successfully loaded .env file")
	}
}
