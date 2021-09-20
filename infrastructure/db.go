package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	fmt.Println(URL)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("database connection estabilished")
	return Database{
		DB: db,
	}

}

// func CloseDB(db *gorm.DB) {
// 	dbcon, err := db.DB()
// 	if err != nil {
// 		panic("unable to close connection")
// 	}
// 	dbcon.Close()
// }
