package server

import (
	"PGL/APIServer/auth"
	"PGL/APIServer/database"
	"PGL/APIServer/log"
	"PGL/APIServer/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//variable for the api version and port used
var (
	apiV = "/api/v1"
	port = ":8080"
)

//Runs the server
func Run() {
	//initialize the database
	database.InitDB()

	//create new multiplexer
	r := mux.NewRouter()
	routes(r)
	r.Use(auth.AuthJWT)

	//secure http listen and serve
	log.Info.Println("APIServer listening at port", port)
	//log.Error.Fatal(http.ListenAndServe(port, r))
	log.Error.Fatal(http.ListenAndServeTLS(port, "certs/cert.pem", "certs/key.pem", r))
}

//success msg
func succMsg(mode, identity string) string {
	msg := fmt.Sprintf("Success %s: %s", mode, identity)
	return msg
}

//other response
func otherRes(w http.ResponseWriter, msg string, status int) {
	var res models.OtherRes
	res.Msg = msg
	if status != 0 {
		w.WriteHeader(status)
	}
	json.NewEncoder(w).Encode(res)
}
