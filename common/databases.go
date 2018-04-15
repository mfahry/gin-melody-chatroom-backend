package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // connect to mysql
)

// DB connection
var DB *gorm.DB

// Init func to open connection
func Init() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Myfiorenta180892@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("db err : ", err)
	}

	db.DB().SetMaxIdleConns(10)
	DB = db

	return DB
}

// GetDB function for get database connection
func GetDB() *gorm.DB {
	return DB
}
