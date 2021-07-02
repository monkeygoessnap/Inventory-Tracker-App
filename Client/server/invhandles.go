package server

import (
	"PGL/Client/api"
	"PGL/Client/models"
	"net/http"
)

//inv handle to display inv
func invs(w http.ResponseWriter, r *http.Request) {
	if !LoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "invs.html", nil)
}

//inv handle to add inv
func addinv(w http.ResponseWriter, r *http.Request) {
	if !LoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		model := "inv"
		//sanitizes the input
		if !inputCheck(r.FormValue("icon"), r.FormValue("name")) {
			http.Error(w, errChars.Error(), http.StatusBadRequest)
			return
		}
		jsonData := models.Inv{
			UserID: getUserID(r),
			Icon:   r.FormValue("icon"),
			Name:   r.FormValue("name"),
		}
		api.Add(model, jsonData)
		http.Redirect(w, r, "/inv", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "addinv.html", nil)
}
