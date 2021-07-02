package server

import (
	"PGL/Client/api"
	"PGL/Client/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//routes
func routes(r *mux.Router) {

	r.HandleFunc("/", index)
	r.HandleFunc("/register", register)
	r.HandleFunc("/login", login)

	r.HandleFunc("/home", home)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/addinv", addinv)
	r.HandleFunc("/additem", additem)
	r.HandleFunc("/profile", profile)
	r.HandleFunc("/editprofile", editprofile)
	r.HandleFunc("/items", items)
	r.HandleFunc("/invs", invs)
	r.HandleFunc("/item/{item}", edititem)

}

//handle to register the user
func register(w http.ResponseWriter, r *http.Request) {
	if LoggedIn(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		if !inputCheck(r.FormValue("username"), r.FormValue("password"), r.FormValue("phone"), r.FormValue("email")) {
			http.Error(w, errChars.Error(), http.StatusBadRequest)
			return
		}
		jsonData := models.User{
			Username: r.FormValue("username"),
			Password: hash(r.FormValue("password")),
			Phone:    r.FormValue("phone"),
			Email:    r.FormValue("email"),
		}
		model := "user"
		data, _ := api.ModelConv(api.Add(model, jsonData), model)
		if user, check := api.ModelConv(api.Get(model, jsonData.Username), model); check {
			updateLogin("add", user.(models.User))
		}
		tpl.ExecuteTemplate(w, "register.html", data.(models.OtherRes).Msg)
		return
	}
	tpl.ExecuteTemplate(w, "register.html", nil)
}

//handle to log the user in and update the prevlogin and currlogin times
func login(w http.ResponseWriter, r *http.Request) {
	if LoggedIn(r) {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		if !inputCheck(r.FormValue("username"), r.FormValue("password")) {
			http.Error(w, errChars.Error(), http.StatusBadRequest)
			return
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		if !checkUser(username, password) {
			data := "Invalid Username/Password"
			tpl.ExecuteTemplate(w, "login.html", data)
			return
		}
		createSession(w, r, username)
		model := "user"
		if user, check := api.ModelConv(api.Get(model, username), model); check {
			updateLogin("edit", user.(models.User))
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", nil)
}

//handle to update the login timings
func updateLogin(mode string, user models.User) {
	model := "setting"
	jsonData := models.UserSetting{
		UserID:    user.ID,
		PrevLogin: user.Setting.CurrLogin,
		CurrLogin: time.Now().Format(time.RFC1123Z),
	}
	switch mode {
	case "add":
		jsonData.PrevLogin = ""
		api.Add(model, jsonData)
	case "edit":
		api.Edit(model, strconv.Itoa(int(user.ID)), jsonData)
	}
}

//handle to log the user out
func logout(w http.ResponseWriter, r *http.Request) {
	if !LoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	removeSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
