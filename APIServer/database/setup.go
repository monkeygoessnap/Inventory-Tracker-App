/*
Package database provides the CRUD functions between the APIServer and the database using
SQL statements.
Uses Gorm to automigrate tables based on the models, as well as issue statements.
*/
package database

import (
	"PGL/APIServer/log"
	"PGL/APIServer/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Provides the db connection pointer
var DB *gorm.DB

//Formats the connection string
func dsnStr() string {

	//get environment variables from env file
	user := os.Getenv("MYSQL_USER")
	pw := os.Getenv("MYSQL_PW")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")

	//formats the connection string
	dsnStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pw, port, dbName)

	return dsnStr
}

//initalize the DB
func InitDB() {

	//connects the database
	var err error
	dsn := dsnStr()
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Error.Fatal(err)
	}
	//auto-migrate & update database tables
	DB.AutoMigrate(&models.User{}, &models.UserSetting{}, &models.Inv{}, &models.Item{},
		&models.Category{})
	log.Info.Println("Successfully loaded database")

}
