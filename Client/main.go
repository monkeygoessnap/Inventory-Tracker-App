/*
Package main runs the program
*/
package main

import (
	"PGL/Client/env"
	"PGL/Client/log"
	"PGL/Client/server"
)

//initializes the key functions
func init() {
	log.InitLog()
	env.InitEnv()
}

//runs the program
func main() {
	log.Info.Println("CLIENT SERVER RUNNING")
	server.Run()
}
