/*
Package server provides the endpoints the for API server,
as well as the mux and https listener
*/
package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//routes
func routes(r *mux.Router) {

	//index
	r.HandleFunc(apiV+"/", index)

	//user routes
	r.HandleFunc(apiV+"/user", getUserID).Methods("GET")
	r.HandleFunc(apiV+"/user", addUser).Methods("POST")
	r.HandleFunc(apiV+"/user/{user}", getUser).Methods("GET")
	r.HandleFunc(apiV+"/user/{user}", delUser).Methods("DELETE")
	r.HandleFunc(apiV+"/user/{user}", editUser).Methods("PUT")

	//inv routes
	r.HandleFunc(apiV+"/inv", addInv).Methods("POST")
	r.HandleFunc(apiV+"/inv/{invid}", getInv).Methods("GET")
	r.HandleFunc(apiV+"/inv/{invid}", delInv).Methods("DELETE")
	r.HandleFunc(apiV+"/inv/{invid}", editInv).Methods("PUT")

	//item routes
	r.HandleFunc(apiV+"/item", getAllItem).Methods("GET")
	r.HandleFunc(apiV+"/item", addItem).Methods("POST")
	r.HandleFunc(apiV+"/item/{itemid}", getItem).Methods("GET")
	r.HandleFunc(apiV+"/item/{itemid}", delItem).Methods("DELETE")
	r.HandleFunc(apiV+"/item/{itemid}", editItem).Methods("PUT")

	//usersetting routes
	r.HandleFunc(apiV+"/setting", getAllSett).Methods("GET")
	r.HandleFunc(apiV+"/setting", addSett).Methods("POST")
	r.HandleFunc(apiV+"/setting/{userid}", getSett).Methods("GET")
	r.HandleFunc(apiV+"/setting/{userid}", delSett).Methods("DELETE")
	r.HandleFunc(apiV+"/setting/{userid}", editSett).Methods("PUT")

	//category routes
	r.HandleFunc(apiV+"/category", getAllCat).Methods("GET")
	r.HandleFunc(apiV+"/category", addCat).Methods("POST")
	r.HandleFunc(apiV+"/category/{catid}", getCat).Methods("GET")
	r.HandleFunc(apiV+"/category/{catid}", delCat).Methods("DELETE")
	r.HandleFunc(apiV+"/category/{catid}", editCat).Methods("PUT")

}

//index page
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to API")
}
