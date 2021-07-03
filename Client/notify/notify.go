/*
Package notify uses Cron to schedule a notification task every hour with the Twilio API
*/
package notify

import (
	"PGL/Client/api"
	"PGL/Client/log"
	"PGL/Client/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron"
)

//Initializes the CRON workers
func InitCron() {
	c := cron.New()
	//Hourly task scheduling
	//@hourly initial setting
	c.AddFunc("@hourly", func() {
		usersid := checkUserToNotify()
		log.Info.Println("Attemping to notify UserID(s) now:", usersid)
		for _, v := range usersid {
			msg := parseMsg(v)
			sendMsg(msg, strconv.Itoa(int(v)))
		}
		log.Info.Println("Cron hourly check, healthy")
	})
	c.Start()
	log.Info.Println("Cron started, running tasks every hourly")
}

//checks which user to notify and saves the id into a slice
func checkUserToNotify() []uint32 {
	model := "setting"
	data, _ := api.ModelAllConv(api.GetAll(model, ""), model)
	var usersToNotify []uint32
	for _, v := range data.([]models.UserSetting) {
		if v.NotiSetting == 2 && checkTime(v) {
			usersToNotify = append(usersToNotify, v.UserID)
		}
	}
	return usersToNotify
}

//checks whether the users setting coincide with the time now of the Cron worker
func checkTime(sett models.UserSetting) bool {
	if len(sett.NotiTime) == 3 {
		sett.NotiTime = "0" + sett.NotiTime
	}
	timenow := time.Now().Format("15") + "00"
	return sett.NotiTime[0:2] == timenow[0:2]
}

//checks whether the item is to be notified
func itemNotify(item models.Item) bool {
	return int(item.Notify)-daysToExpire(item.Expiry) == 0
}

//checks how many day it is to expiry
func daysToExpire(unix uint64) int {
	t1 := time.Now()
	t2 := time.Unix(int64(unix), 0)
	y, m, d := t2.Date()
	u2 := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = t1.In(t2.Location()).Date()
	u1 := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	days := u2.Sub(u1) / (24 * time.Hour)
	if int(days) < 0 {
		return -1
	} else {
		return int(days)
	}
}

//reformats the unix to date string for stringmsg
func toDate(unix uint64) string {
	return time.Unix(int64(unix), 0).Format("02/01/2006")
}

//parses a msg for each of the users in the range
func parseMsg(userid uint32) string {
	model := "item"
	var msg []string
	var retmsg string
	data, _ := api.ModelAllConv(api.GetAll(model, strconv.Itoa(int(userid))), model)
	for _, v := range data.([]models.Item) {
		if itemNotify(v) {
			itemmsg := fmt.Sprintf("%s is expiring in %v day(s) on %s !", v.Name, v.Notify, toDate(v.Expiry))
			msg = append(msg, itemmsg)
		}
	}
	retmsg = strings.Join(msg, " | ")
	return retmsg
}

//msg to tell the twilio api to send a msg to the fella
func sendMsg(msg string, userid string) {
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	model := "user"
	info, _ := api.ModelConv(api.GetAll(model, userid), model)
	//get fella's phone number from setting
	numberTo := "+65" + info.(models.User).Phone
	//get twilio's generated phone number
	numberFrom := os.Getenv("TWILIO_NO")
	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", numberTo)
	msgData.Set("From", numberFrom)
	msgData.Set("Body", msg)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			log.Info.Println("twilio success", data["sid"], "UserID:", userid)
		}
	} else {
		log.Info.Println("twilio error", resp.Status, "UserID", userid)
	}
}
