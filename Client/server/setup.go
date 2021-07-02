/*
Package server provides the handles and mux to run the server
*/
package server

import (
	"PGL/Client/log"
	"PGL/Client/notify"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//variables for port and tpl pointer
var (
	port = ":8081"
	tpl  *template.Template
)

//runs the server
func Run() {

	//parses the template with the custom functions
	tpl = template.Must(template.New("").Funcs(functionMap).ParseGlob("templates/*.html"))
	//create new multiplexer
	r := mux.NewRouter()
	//routes
	routes(r)
	//templating directory
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	//go function to run the cron concurrently
	go notify.InitCron()
	//secure http listen and serve
	log.Info.Println("Client listening at port", port)
	log.Error.Fatal(http.ListenAndServeTLS(port, "certs/cert.pem", "certs/key.pem", r))
}

//custom funcs for the templates
var functionMap = template.FuncMap{
	"toDate": func(unix uint64) string {
		if unix == 0 {
			return "Nil"
		}
		return time.Unix(int64(unix), 0).Format("02/01/2006")
	},
	"toDay": func(unix uint64) string {
		if unix == 0 {
			return ""
		}
		t1 := time.Now()
		t2 := time.Unix(int64(unix), 0)
		y, m, d := t2.Date()
		u2 := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
		y, m, d = t1.In(t2.Location()).Date()
		u1 := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
		days := u2.Sub(u1) / (24 * time.Hour)
		ret := strconv.Itoa(int(days))
		if int(days) < 0 {
			return "Expired"
		} else {
			return ret + " DaysToGo"
		}
	},
	"toDate2": func(unix uint64) string {
		if unix == 0 {
			return "Nil"
		}
		return time.Unix(int64(unix), 0).Format("2006-01-02")
	},
}
