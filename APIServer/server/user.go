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

//get user based on the id
func getUserID(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	userid, _ := strconv.ParseUint(v["userid"][0], 10, 32)
	info, err := database.GetUserID(uint32(userid))
	if err != nil {
		otherRes(w, err.Error(), http.StatusInternalServerError)
	}
	log.Info.Println(succMsg("GET USER BY ID", v["userid"][0]))
	json.NewEncoder(w).Encode(info)
}

//get user based on username
func getUser(w http.ResponseWriter, r *http.Request) {

	//get params
	params := mux.Vars(r)
	username := params["user"]
	//get data from database
	userInfo, err := database.GetUser(username)
	//error handling
	switch err {
	case database.ErrNotFound:
		otherRes(w, err.Error(), http.StatusNotFound)
	case database.ErrInternal:
		otherRes(w, err.Error(), http.StatusInternalServerError)
	default:
		log.Info.Println(succMsg("GET USER", username))
		json.NewEncoder(w).Encode(userInfo)
	}
}

//add user
func addUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" && r.Header.Get("Content-Type") == "application/json" {
		var info models.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &info)
		err := database.AddUser(info)
		switch err {
		case database.ErrUserTaken:
			otherRes(w, err.Error(), http.StatusConflict)
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		default:
			log.Info.Println(succMsg("ADD USER", info.Username))
			otherRes(w, "success", http.StatusAccepted)
		}
	}
}

//delete user
func delUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		params := mux.Vars(r)
		username := params["user"]
		err := database.DelUser(username)
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		case database.ErrNotFound:
			otherRes(w, err.Error(), http.StatusNotFound)
		default:
			log.Info.Println(succMsg("DEL USER", username))
			otherRes(w, "success", http.StatusAccepted)
		}
	}
}

//edit user
func editUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" && r.Header.Get("Content-Type") == "application/json" {
		var newInfo models.User
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &newInfo)
		params := mux.Vars(r)
		username := params["user"]
		err := database.EditUser(newInfo, username)
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		default:
			log.Info.Println(succMsg("UPDATE USER", username))
			otherRes(w, "success", http.StatusAccepted)
		}
	}

}
