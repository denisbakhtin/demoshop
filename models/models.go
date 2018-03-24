package models

import (
	"log"

	"github.com/jinzhu/gorm"
	//required by gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//InitDB establishes connection to database and saves its handler into db *sqlx.DB
func InitDB(connection string) {
	var err error
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	//automigrate
	db.AutoMigrate(&Product{}, &Vendor{})
	seed() //seed db
}

//GetDB returns database handler
func GetDB() *gorm.DB {
	return db
}
