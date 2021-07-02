package database

import (
	"PGL/APIServer/log"
	"PGL/APIServer/models"

	"gorm.io/gorm"
)

//Get all categories based on the user id
func GetAllCat(userid uint32) ([]models.Category, error) {
	var cats []models.Category
	err := DB.Where("user_id = ?", userid).Find(&cats).Error
	if err != nil {
		log.Warning.Println(err)
		return nil, ErrInternal
	}
	return cats, nil
}

//Get category based on the category id
func GetCat(catid uint32) (models.Category, error) {
	var catInfo models.Category

	err := DB.Where("id = ?", catid).First(&catInfo).Error

	if err == gorm.ErrRecordNotFound {
		log.Info.Println(err)
		return catInfo, ErrNotFound
	} else if err != nil {
		log.Warning.Println(err)
		return catInfo, ErrInternal
	}
	return catInfo, nil
}

//Edits the category based on the category id
func EditCat(catInfo models.Category, catid uint32) error {

	err := DB.Where("id = ?", catid).Updates(catInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Adds the category based on the category id
func AddCat(catInfo models.Category) error {

	err := DB.Create(&catInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Deletes the category based on the category id
func DelCat(catid uint32) error {

	result := DB.Where("id = ?", catid).Delete(&models.Category{})
	if result.Error != nil {
		log.Warning.Println(result.Error)
		return ErrInternal
	} else if result.RowsAffected < 1 {
		return ErrNotFound
	}
	return nil
}
