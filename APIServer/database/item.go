package database

import (
	"PGL/APIServer/log"
	"PGL/APIServer/models"

	"gorm.io/gorm"
)

//Gets all items based on the user id
func GetAllItem(userid uint32) ([]models.Item, error) {
	var items []models.Item

	err := DB.Where("inv_id IN (SELECT id FROM invs WHERE user_id = ?)", userid).Find(&items).Error
	if err != nil {
		log.Warning.Println(err)
		return nil, ErrInternal
	}
	return items, nil
}

//Get item based on the itemid
func GetItem(itemid uint32) (models.Item, error) {
	var itemInfo models.Item

	err := DB.Where("id = ?", itemid).First(&itemInfo).Error

	if err == gorm.ErrRecordNotFound {
		log.Info.Println(err)
		return itemInfo, ErrNotFound
	} else if err != nil {
		log.Warning.Println(err)
		return itemInfo, ErrInternal
	}
	return itemInfo, nil
}

//Edit item based on the itemid
func EditItem(itemInfo models.Item, itemid uint32) error {

	err := DB.Where("id = ?", itemid).Updates(itemInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Add item
func AddItem(itemInfo models.Item) error {

	err := DB.Create(&itemInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Deletes the item based on the item id
func DelItem(itemid uint32) error {

	result := DB.Where("id = ?", itemid).Delete(&models.Item{})
	if result.Error != nil {
		log.Warning.Println(result.Error)
		return ErrInternal
	} else if result.RowsAffected < 1 {
		return ErrNotFound
	}
	return nil
}
