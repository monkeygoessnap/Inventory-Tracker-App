package server

import (
	"PGL/Client/api"
	"PGL/Client/models"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//handle to edit the item
func edititem(w http.ResponseWriter, r *http.Request) {
	if !LoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//get the params itemid
	params := mux.Vars(r)
	itemid := params["item"]
	modelu := "user"
	username := getUsername(r)
	user, _ := api.ModelConv(api.Get(modelu, username), modelu)
	modeli := "item"
	itemInfo, _ := api.ModelConv(api.Get(modeli, itemid), modeli)
	if r.Method == http.MethodPost {
		//sanitizes the input
		if !inputCheck(r.FormValue("icon"), r.FormValue("name"), r.FormValue("category"), r.FormValue("storage"),
			r.FormValue("expiry"), r.FormValue("notify")) {
			http.Error(w, errChars.Error(), http.StatusBadRequest)
			return
		}
		modeli := "item"
		var invid uint32
		//creates the inventory if it doesn't exist
		for _, v := range user.(models.User).Inv {
			if v.Name == r.FormValue("inv") {
				invid = v.ID
			}
		}
		if invid == 0 {
			newInv := models.Inv{
				UserID: getUserID(r),
				Name:   r.FormValue("inv"),
			}
			modelinv := "inv"
			api.Add(modelinv, newInv)
		}
		//appends the new invid to the created item
		data, _ := api.ModelConv(api.Get(modelu, username), modelu)
		for _, v := range data.(models.User).Inv {
			if v.Name == r.FormValue("inv") {
				invid = v.ID
			}
		}
		notify, _ := strconv.Atoi(r.FormValue("notify"))
		expiry, _ := time.Parse("2006-01-02", r.FormValue("expiry"))
		unixd := uint64(expiry.Unix())
		// if r.FormValue("notify") == "" {
		// 	unixd = 0
		// }
		var idle uint64
		if r.FormValue("idle") != "" {
			idle = uint64(time.Now().Unix())
		}
		jsonData := models.Item{
			InvID:    uint32(invid),
			Icon:     r.FormValue("icon"),
			Name:     r.FormValue("name"),
			Category: r.FormValue("category"),
			Storage:  r.FormValue("storage"),
			Expiry:   unixd,
			Idle:     idle,
			Notify:   uint32(notify),
		}
		api.Edit(modeli, itemid, jsonData)
		http.Redirect(w, r, "/items", http.StatusSeeOther)
		return
	}
	var invname string
	for _, v := range user.(models.User).Inv {
		if v.ID == itemInfo.(models.Item).InvID {
			invname = v.Name
		}
	}
	var data []interface{}
	data = append(data, itemInfo.(models.Item), invname)
	tpl.ExecuteTemplate(w, "edititem.html", data)
}

//handle to show the items
func items(w http.ResponseWriter, r *http.Request) {
	if !LoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		if !isNum(r.FormValue("delete"), r.FormValue("edit")) {
			http.Error(w, errChars.Error(), http.StatusBadRequest)
			return
		}
		deleteid := r.FormValue("delete")
		editid := r.FormValue("edit")
		if deleteid != "" {
			model := "item"
			api.Del(model, deleteid)
		}
		if editid != "" {
			http.Redirect(w, r, "/item/"+editid, http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/items", http.StatusSeeOther)
		return
	}
	model := "item"
	userid := strconv.Itoa(int(getUserID(r)))
	data, _ := api.ModelAllConv(api.GetAll(model, userid), model)
	sort.SliceStable(data.([]models.Item), func(i, j int) bool {
		return int(data.([]models.Item)[i].Expiry) < int(data.([]models.Item)[j].Expiry)
	})
	tpl.ExecuteTemplate(w, "items.html", data.([]models.Item))
}

//handle to add the item
func additem(w http.ResponseWriter, r *http.Request) {
	if !LoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	modelu := "user"
	username := getUsername(r)
	user, _ := api.ModelConv(api.Get(modelu, username), modelu)
	if r.Method == http.MethodPost {
		//sanitizes the input
		if !inputCheck(r.FormValue("icon"), r.FormValue("name"), r.FormValue("category"), r.FormValue("storage"),
			r.FormValue("expiry"), r.FormValue("notify")) {
			http.Error(w, errChars.Error(), http.StatusBadRequest)
			return
		}
		modeli := "item"
		var invid uint32
		//creates the inventory if it doesn't exist
		for _, v := range user.(models.User).Inv {
			if v.Name == r.FormValue("inv") {
				invid = v.ID
			}
		}
		if invid == 0 {
			newInv := models.Inv{
				UserID: getUserID(r),
				Name:   r.FormValue("inv"),
			}
			modelinv := "inv"
			api.Add(modelinv, newInv)
		}
		//appends the new invid to the created item
		data, _ := api.ModelConv(api.Get(modelu, username), modelu)
		for _, v := range data.(models.User).Inv {
			if v.Name == r.FormValue("inv") {
				invid = v.ID
			}
		}
		notify, _ := strconv.Atoi(r.FormValue("notify"))
		expiry, _ := time.Parse("2006-01-02", r.FormValue("expiry"))
		unixd := uint64(expiry.Unix())
		// if r.FormValue("notify") == "" {
		// 	unixd = 0
		// }
		var idle uint64
		if r.FormValue("idle") != "" {
			idle = uint64(time.Now().Unix())
		}
		jsonData := models.Item{
			InvID:    uint32(invid),
			Icon:     r.FormValue("icon"),
			Name:     r.FormValue("name"),
			Category: r.FormValue("category"),
			Storage:  r.FormValue("storage"),
			Expiry:   unixd,
			Idle:     idle,
			Notify:   uint32(notify),
		}
		api.Add(modeli, jsonData)
		http.Redirect(w, r, "/items", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "additem.html", user.(models.User))
}
