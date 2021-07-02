package database

import (
	"PGL/APIServer/log"
	"PGL/APIServer/models"
)

//Get all settings based on the id
func GetAllSett() ([]models.UserSetting, error) {
	var settings []models.UserSetting
	err := DB.Find(&settings).Error
	if err != nil {
		log.Warning.Println(err)
		return nil, ErrInternal
	}
	return settings, nil
}

//Get the setting based on the userid
func GetSett(userid uint32) (models.UserSetting, error) {
	var setting models.UserSetting
	err := DB.Where("user_id = ?", userid).Find(&setting).Error
	if err != nil {
		log.Warning.Println(err)
		return setting, err
	}
	return setting, nil
}

//Edit the setting based on the userid
func EditSett(settInfo models.UserSetting, userid uint32) error {

	err := DB.Where("user_id = ?", userid).Updates(settInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Add a setting
func AddSett(settInfo models.UserSetting) error {

	err := DB.Create(&settInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Deletes the setting based on the userid
func DelSett(userid uint32) error {

	result := DB.Where("user_id = ?", userid).Delete(&models.UserSetting{})
	if result.Error != nil {
		log.Warning.Println(result.Error)
		return ErrInternal
	} else if result.RowsAffected < 1 {
		return ErrNotFound
	}
	return nil
}
