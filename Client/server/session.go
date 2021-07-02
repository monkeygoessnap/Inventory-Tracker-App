package server

import (
	"PGL/Client/api"
	"PGL/Client/log"
	"PGL/Client/models"
	"net/http"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

//sync variables for the package
var (
	//mutex sync.Mutex
	wg sync.WaitGroup
)

//session map
var mapSessions = map[string]SessionInfo{}

//session struct
type SessionInfo struct {
	UserID   uint32
	Username string
}

//function to create a new session
func createSession(w http.ResponseWriter, r *http.Request, username string) {
	//new random uuid using satori
	id := uuid.NewV4()
	myCookie := &http.Cookie{
		Name:     "myCookie",
		Value:    id.String(),
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	}
	//issues new cookie
	http.SetCookie(w, myCookie)
	//maps a new session
	var newSession SessionInfo
	model := "user"
	if user, check := api.ModelConv(api.Get(model, username), model); check {
		newSession.UserID = user.(models.User).ID
		newSession.Username = username
	}
	mapSessions[myCookie.Value] = newSession
}

//function to remove an existing session
func removeSession(w http.ResponseWriter, r *http.Request) {
	myCookie, _ := r.Cookie("myCookie")
	//deletes the mapped session
	delete(mapSessions, myCookie.Value)
	//removes the cookie
	myCookie = &http.Cookie{
		Name:   "myCookie",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, myCookie)
}

//function to check whether a user is logged in by requesting
//and checking the details from the cookies against the session map
func LoggedIn(r *http.Request) bool {
	if _, err := r.Cookie("myCookie"); err == nil {
		return len(getUsername(r)) > 0
	}
	return false
}

//function to get the username from the mapped session
func getUsername(r *http.Request) string {
	// get current session cookie
	myCookie, _ := r.Cookie("myCookie")
	// if the user exists already, get user
	if info, ok := mapSessions[myCookie.Value]; ok {
		return info.Username
	}
	return ""
}

//function to get userid
func getUserID(r *http.Request) uint32 {
	// get current session cookie
	myCookie, _ := r.Cookie("myCookie")
	// if the user exists already, get user
	if info, ok := mapSessions[myCookie.Value]; ok {
		return info.UserID
	}
	return 0
}

//function to check whether a user credentials are correct
func checkUser(username, password string) bool {
	//var user models.Users
	model := "user"
	res, found := api.ModelConv(api.Get(model, username), model)
	if !found {
		return false
	}
	return checkHash(res.(models.User).Password, password)
}

//function to hash password
func hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hash)
}

//function to check the hashed password against the input
func checkHash(hashedpw, pw string) bool {
	//timing attack prevention
	wg.Add(1)
	go func() {
		time.Sleep(50 * time.Millisecond)
		wg.Done()
	}()
	defer wg.Wait()
	err := bcrypt.CompareHashAndPassword([]byte(hashedpw), []byte(pw))
	if err != nil {
		log.Info.Println(err)
		return false
	}
	return true
}
