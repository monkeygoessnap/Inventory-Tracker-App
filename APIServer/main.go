/*
Package main to run the APIServer program
*/
package main

import (
	"PGL/APIServer/env"
	"PGL/APIServer/log"
	"PGL/APIServer/server"
)

//initialize the env file
func init() {

	//init log & check
	log.InitLog()
	//init env file & check
	env.InitEnv()

}

//runs the program
func main() {

	log.Info.Println("API SERVER RUNNING")
	server.Run()

}
