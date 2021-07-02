package server

import (
	"PGL/Client/api"
	"PGL/Client/models"
	"net/http"
	"strconv"
)

//handle for userhome
func home(w http.ResponseWriter, r *http.Request) {
	if !LoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	model := "user"
	username := getUsername(r)
	data, _ := api.ModelConv(api.Get(model, username), model)
	tpl.ExecuteTemplate(w, "home.html", data)
}

//handle for user profile
func profile(w http.ResponseWriter, r *http.Request) {
	if !LoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	model := "user"
	user, _ := api.ModelConv(api.Get(model, getUsername(r)), model)
	var noti string
	switch user.(models.User).Setting.NotiSetting {
	case 1:
		noti = "Do not notify"
	case 2:
		noti = "Notify"
	}
	var data []interface{}
	data = append(data, user.(models.User), noti)
	tpl.ExecuteTemplate(w, "profile.html", data)
}

//handle for editing user profile
func editprofile(w http.ResponseWriter, r *http.Request) {
	if !LoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		//sanitizes input
		if !inputCheck(r.FormValue("phone"), r.FormValue("email")) {
			http.Error(w, errChars.Error(), http.StatusBadRequest)
			return
		}
		username := getUsername(r)
		userid := getUserID(r)
		jsonData1 := models.User{
			Phone: r.FormValue("phone"),
			Email: r.FormValue("email"),
		}
		notisetting, _ := strconv.Atoi(r.FormValue("setting"))
		jsonData2 := models.UserSetting{
			NotiTime:    r.FormValue("notitime"),
			NotiSetting: uint32(notisetting),
		}
		model1 := "user"
		model2 := "setting"
		api.Edit(model1, username, jsonData1)
		api.Edit(model2, strconv.Itoa(int(userid)), jsonData2)
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}
	model := "user"
	user, _ := api.ModelConv(api.Get(model, getUsername(r)), model)
	tpl.ExecuteTemplate(w, "editprofile.html", user.(models.User))
}

//index
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "index.html", nil)
}
