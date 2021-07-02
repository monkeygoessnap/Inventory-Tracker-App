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

//Get all categories
func getAllCat(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	userid, _ := strconv.ParseUint(v["userid"][0], 10, 32)
	info, err := database.GetAllCat(uint32(userid))
	if err != nil {
		otherRes(w, err.Error(), http.StatusInternalServerError)
	}
	log.Info.Println(succMsg("GET ALL CATEGORY", v["userid"][0]))
	json.NewEncoder(w).Encode(info)
}

//get category
func getCat(w http.ResponseWriter, r *http.Request) {

	//get params
	params := mux.Vars(r)
	catid, _ := strconv.ParseUint(params["catid"], 10, 32)
	//get data from database
	info, err := database.GetCat(uint32(catid))
	//error handling
	switch err {
	case database.ErrNotFound:
		otherRes(w, err.Error(), http.StatusNotFound)
	case database.ErrInternal:
		otherRes(w, err.Error(), http.StatusInternalServerError)
	default:
		log.Info.Println(succMsg("GET CATEGORY", params["catid"]))
		json.NewEncoder(w).Encode(info)
	}
}

//add category
func addCat(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" && r.Header.Get("Content-Type") == "application/json" {
		var info models.Category
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &info)
		err := database.AddCat(info)
		if err == database.ErrInternal {
			otherRes(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Info.Println(succMsg("ADD CATEGORY", strconv.Itoa(int(info.UserID))))
		otherRes(w, "success", http.StatusAccepted)
	}
}

//delete category
func delCat(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		params := mux.Vars(r)
		catid, _ := strconv.ParseUint(params["catid"], 10, 32)
		err := database.DelCat(uint32(catid))
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		case database.ErrNotFound:
			otherRes(w, err.Error(), http.StatusNotFound)
		default:
			log.Info.Println(succMsg("DEL CATEGORY", params["catid"]))
			otherRes(w, "success", http.StatusAccepted)
		}
	}
}

//edit category
func editCat(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" && r.Header.Get("Content-Type") == "application/json" {
		var info models.Category
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &info)
		params := mux.Vars(r)
		catid, _ := strconv.ParseUint(params["catid"], 10, 32)
		err := database.EditCat(info, uint32(catid))
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		default:
			log.Info.Println(succMsg("UPDATE CATEGORY", params["catid"]))
			otherRes(w, "success", http.StatusAccepted)
		}
	}

}
