package database

import (
	"PGL/APIServer/log"
	"PGL/APIServer/models"

	"gorm.io/gorm"
)

//Get inventory records based on the inv id
func GetInv(invid uint32) (models.Inv, error) {
	var invInfo models.Inv

	err := DB.Where("id = ?", invid).First(&invInfo).Error

	if err == gorm.ErrRecordNotFound {
		log.Info.Println(err)
		return invInfo, ErrNotFound
	} else if err != nil {
		log.Warning.Println(err)
		return invInfo, ErrInternal
	}

	invInfo.Items, err = getItems(invid)
	if err != nil {
		return invInfo, ErrInternal
	}

	return invInfo, nil
}

//Edit inventory record based on the inv id
func EditInv(invInfo models.Inv, invid uint32) error {

	err := DB.Where("id = ?", invid).Updates(invInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Add inv record based on the inv id
func AddInv(invInfo models.Inv) error {

	err := DB.Create(&invInfo).Error
	if err != nil {
		log.Warning.Println(err)
		return ErrInternal
	}
	return nil
}

//Deletes the inv record based on the inv id
func DelInv(invid uint32) error {

	result := DB.Where("id = ?", invid).Delete(&models.Inv{})
	if result.Error != nil {
		log.Warning.Println(result.Error)
		return ErrInternal
	} else if result.RowsAffected < 1 {
		return ErrNotFound
	}
	return nil
}
