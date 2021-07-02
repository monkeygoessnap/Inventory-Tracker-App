package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"PGL/APIServer/database"
	"PGL/APIServer/log"
	"PGL/APIServer/models"

	"github.com/gorilla/mux"
)

//get inventory
func getInv(w http.ResponseWriter, r *http.Request) {

	//get params
	params := mux.Vars(r)
	invid, _ := strconv.ParseUint(params["invid"], 10, 32)
	//get data from database
	info, err := database.GetInv(uint32(invid))
	//error handling
	switch err {
	case database.ErrNotFound:
		otherRes(w, err.Error(), http.StatusNotFound)
	case database.ErrInternal:
		otherRes(w, err.Error(), http.StatusInternalServerError)
	default:
		log.Info.Println(succMsg("GET INV", params["invid"]))
		json.NewEncoder(w).Encode(info)
	}
}

//add inventory
func addInv(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" && r.Header.Get("Content-Type") == "application/json" {
		var info models.Inv
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &info)
		err := database.AddInv(info)
		if err == database.ErrInternal {
			otherRes(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Info.Println(succMsg("ADD INV", info.Name))
		otherRes(w, "success", http.StatusAccepted)
	}
}

//delete inventory
func delInv(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		params := mux.Vars(r)
		invid, _ := strconv.ParseUint(params["invid"], 10, 32)
		err := database.DelInv(uint32(invid))
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		case database.ErrNotFound:
			otherRes(w, err.Error(), http.StatusNotFound)
		default:
			log.Info.Println(succMsg("DEL INV", params["invid"]))
			otherRes(w, "success", http.StatusAccepted)
		}
	}
}

//edit inventory
func editInv(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" && r.Header.Get("Content-Type") == "application/json" {
		var info models.Inv
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &info)
		params := mux.Vars(r)
		invid, _ := strconv.ParseUint(params["invid"], 10, 32)
		err := database.EditInv(info, uint32(invid))
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		default:
			log.Info.Println(succMsg("UPDATE INV", params["invid"]))
			otherRes(w, "success", http.StatusAccepted)
		}
	}

}
