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

//get all items
func getAllItem(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	userid, _ := strconv.ParseUint(v["userid"][0], 10, 32)
	info, err := database.GetAllItem(uint32(userid))
	if err != nil {
		otherRes(w, err.Error(), http.StatusInternalServerError)
	}
	log.Info.Println(succMsg("GET ALL ITEM", v["userid"][0]))
	json.NewEncoder(w).Encode(info)
}

//get item
func getItem(w http.ResponseWriter, r *http.Request) {

	//get params
	params := mux.Vars(r)
	itemid, _ := strconv.ParseUint(params["itemid"], 10, 32)
	//get data from database
	info, err := database.GetItem(uint32(itemid))
	//error handling
	switch err {
	case database.ErrNotFound:
		otherRes(w, err.Error(), http.StatusNotFound)
	case database.ErrInternal:
		otherRes(w, err.Error(), http.StatusInternalServerError)
	default:
		log.Info.Println(succMsg("GET ITEM", params["itemid"]))
		json.NewEncoder(w).Encode(info)
	}
}

//add item
func addItem(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" && r.Header.Get("Content-Type") == "application/json" {
		var info models.Item
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &info)
		err := database.AddItem(info)
		if err == database.ErrInternal {
			otherRes(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Info.Println(succMsg("ADD Item", info.Name))
		otherRes(w, "success", http.StatusAccepted)
	}
}

//detele item
func delItem(w http.ResponseWriter, r *http.Request) {

	if r.Method == "DELETE" {
		params := mux.Vars(r)
		itemid, _ := strconv.ParseUint(params["itemid"], 10, 32)
		err := database.DelItem(uint32(itemid))
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		case database.ErrNotFound:
			otherRes(w, err.Error(), http.StatusNotFound)
		default:
			log.Info.Println(succMsg("DEL ITEM", params["itemid"]))
			otherRes(w, "success", http.StatusAccepted)
		}
	}
}

//edit item
func editItem(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" && r.Header.Get("Content-Type") == "application/json" {
		var info models.Item
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &info)
		params := mux.Vars(r)
		itemid, _ := strconv.ParseUint(params["itemid"], 10, 32)
		err := database.EditItem(info, uint32(itemid))
		switch err {
		case database.ErrInternal:
			otherRes(w, err.Error(), http.StatusInternalServerError)
		default:
			log.Info.Println(succMsg("UPDATE ITEM", params["itemid"]))
			otherRes(w, "success", http.StatusAccepted)
		}
	}

}
