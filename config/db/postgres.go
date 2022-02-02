package db

import (
	"github.com/jinzhu/gorm"
	"log"
)

func Connect(postgresUrl string) *gorm.DB {

	log.Println("postgresUrl : ", postgresUrl)

	db, err := gorm.Open("postgres", postgresUrl)
	if err != nil {
		log.Println("Error while connecting to repository: ", err)
		panic(err)
	}

	return db
}
