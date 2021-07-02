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

//get all setting
func getAllSett(w http.ResponseWriter, r *http.Request) {
	info, err := database.GetAllSett()
	switch err {
	case database.ErrInternal:
		otherRes(w, err.Error(), http.StatusInternalServerError)
	default:
		log.Info.Println(succMsg("GETALL USERSETTING", ""))
		json.NewEncoder(w).Encode(info)
	}
}

//get setting
func getSett(w http.ResponseWriter, r *http.Request) {

	//get params
	params := mux.Vars(r)
	userid, _ := strconv.ParseUint(params["userid"], 10, 32)
	//get data from database
	info, err := database.GetSett(uint32(userid))
	//error handling
	switch err {
	case database.ErrNotFound:
		otherRes(w, err.Error(), http.StatusNotFound)
	case database.ErrInternal:
		otherRes(w, err.Error(), http.StatusInternalServerError)
	default:
		log.Info.Println(succMsg("GET USERSETTING", params["userid"]))
		json.NewEncoder(w).Encode(info)
	}
}

//add setting
func addSett(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" && r.Header.Get("Content-Type") == "application/json" {
		var info models.UserSetting
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &info)
		err := database.AddSett(info)
		if err == database.ErrInternal {
			otherRes(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Info.Println(succMsg("ADD USERSETTING", strconv.Itoa(int(info.UserID))))
		otherRes(w, "success", http.StatusAccepted)
	}
}

//delete setting
func delSett(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		params := mux.Vars(r)
		userid, _ := strconv.ParseUint(params["userid"], 10, 32)
		err := database.DelSett(uint32(userid))
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		case database.ErrNotFound:
			otherRes(w, err.Error(), http.StatusNotFound)
		default:
			log.Info.Println(succMsg("DEL USERSETTING", params["userid"]))
			otherRes(w, "success", http.StatusAccepted)
		}
	}
}

//edit setting
func editSett(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" && r.Header.Get("Content-Type") == "application/json" {
		var info models.UserSetting
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &info)
		params := mux.Vars(r)
		userid, _ := strconv.ParseUint(params["userid"], 10, 32)
		err := database.EditSett(info, uint32(userid))
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		default:
			log.Info.Println(succMsg("UPDATE USERSETTING", params["userid"]))
			otherRes(w, "success", http.StatusAccepted)
		}
	}

}
