package databases

import (
	"assignment-2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "hacktiv"
	password = "koinworks12"
	dbname   = "assignment2db"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	//fmt.Println(config)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Connecting to database is Error: ", err)
	}
	db.AutoMigrate(models.Order{}, models.Item{})

}

func GetDB() *gorm.DB {
	return db
}
