package database

import (
	"fmt"
	"log"
	"sesi7-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "esra"
	password = "koinworks"
	db_name  = "hacktiv_sesi7"
)

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, db_name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Default().Println("Connection db success")
	//db.Debug().AutoMigrate(models.User{}, models.Product{})
	err = migration(db)
	if err != nil {
		panic(err)
	}
	return db
}
func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(models.User{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(models.Product{}); err != nil {
		return err
	}
	return nil
}