package api

import (
	"PGL/Client/models"
	"encoding/json"
)

//Converts the JSON response to appropriate interface models to be asserted later on
func ModelConv(data []byte, model string) (interface{}, bool) {

	//check if incoming data is a msg
	var otherRes models.OtherRes
	json.Unmarshal(data, &otherRes)
	if otherRes.Msg != "" {
		return otherRes, false
	}
	switch model {
	case "user":
		var user models.User
		json.Unmarshal(data, &user)
		return user, true
	case "inv":
		var inv models.Inv
		json.Unmarshal(data, &inv)
		return inv, true
	case "item":
		var item models.Item
		json.Unmarshal(data, &item)
		return item, true
	case "setting":
		var sett models.UserSetting
		json.Unmarshal(data, &sett)
		return sett, true
	case "category":
		var cat models.Category
		json.Unmarshal(data, &cat)
		return cat, true
	default:
		return otherRes, false
	}
}

//Converts the JSON response to appropriate interface models to be asserted later on, for multiple responses
func ModelAllConv(data []byte, model string) (interface{}, bool) {

	//check if incoming data is a msg
	var otherRes models.OtherRes
	json.Unmarshal(data, &otherRes)
	if otherRes.Msg != "" {
		return otherRes, false
	}
	switch model {
	case "user":
		var user []models.User
		json.Unmarshal(data, &user)
		return user, true
	case "inv":
		var inv []models.Inv
		json.Unmarshal(data, &inv)
		return inv, true
	case "item":
		var item []models.Item
		json.Unmarshal(data, &item)
		return item, true
	case "setting":
		var sett []models.UserSetting
		json.Unmarshal(data, &sett)
		return sett, true
	case "category":
		var cat []models.Category
		json.Unmarshal(data, &cat)
		return cat, true
	default:
		return otherRes, false
	}
}
