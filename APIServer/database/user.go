package database

import (
	"PGL/APIServer/log"
	"PGL/APIServer/models"

	"gorm.io/gorm"
)

//Gets User record based on the id
func GetUserID(userid uint32) (models.User, error) {
	var userInfo models.User

	err := DB.Where("id = ?", userid).First(&userInfo).Error

	if err == gorm.ErrRecordNotFound {
		log.Info.Println(err)
		return userInfo, ErrNotFound
	} else if err != nil {
		log.Warning.Println(err)
		return userInfo, ErrInternal
	}
	userInfo.Setting, err = getSettings(userInfo.ID)
	if err != nil {
		return userInfo, ErrInternal
	}
	return userInfo, nil
}

//Gets user record based on the username
func GetUser(username string) (models.User, error) {
	var userInfo models.User

	err := DB.Where("username = ?", username).First(&userInfo).Error

	if err == gorm.ErrRecordNotFound {
		log.Info.Println(err)
		return userInfo, ErrNotFound
	} else if err != nil {
		log.Warning.Println(err)
		return userInfo, ErrInternal
	}
	userInfo.Inv, err = getInvs(userInfo.ID)
	if err != nil {
		return userInfo, ErrInternal
	}
	userInfo.Setting, err = getSettings(userInfo.ID)
	if err != nil {
		return userInfo, ErrInternal
	}
	return userInfo, nil
}

//Gets inv records to populate in user struct
func getInvs(userid uint32) ([]models.Inv, error) {
	var inv models.Inv
	var invs []models.Inv
	rows, err := DB.Model(&models.Inv{}).Where("user_id = ?", userid).Rows()
	defer rows.Close()
	if err != nil {
		log.Warning.Println(err)
		return nil, err
	}
	for rows.Next() {
		DB.ScanRows(rows, &inv)
		inv.Items, err = getItems(inv.ID)
		if err != nil {
			return nil, err
		}
		invs = append(invs, inv)
	}
	return invs, nil
}

//Gets items to populate in the inventory
func getItems(invid uint32) ([]models.Item, error) {
	var items []models.Item
	err := DB.Where("inv_id = ?", invid).Find(&items).Error
	if err != nil {
		log.Warning.Println(err)
		return nil, err
	}
	return items, nil
}

//Gets usersettings to populate in the user struct
func getSettings(userid uint32) (models.UserSetting, error) {
	var setting models.UserSetting
	err := DB.Where("user_id = ?", userid).Find(&setting).Error
	if err != nil {
		log.Warning.Println(err)
		return setting, err
	}
	return setting, nil
}

//Edits user record based on the username
func EditUser(userInfo models.User, username string) error {

	err := DB.Where("username = ?", username).Updates(userInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Adds user
func AddUser(userInfo models.User) error {

	err := DB.Where("username = ?", userInfo.Username).First(&models.User{}).Error
	if err != gorm.ErrRecordNotFound {
		log.Info.Println(ErrUserTaken)
		return ErrUserTaken
	}
	err = DB.Create(&userInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Deletes user based on username
func DelUser(username string) error {

	result := DB.Where("username = ?", username).Delete(&models.User{})
	if result.Error != nil {
		log.Warning.Println(result.Error)
		return ErrInternal
	} else if result.RowsAffected < 1 {
		return ErrNotFound
	}
	return nil
}
