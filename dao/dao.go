package dao

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var db *gorm.DB

//var cfg *ini.File

func init() {

	var database = &Database{
		Type:     "mysql",
		User:     "root",
		Password: "123",
		Host:     "localhost",
		Name:     "official",
		//		TablePrefix string
	}

	var err error
	db, err = gorm.Open(
		database.Type,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			database.User,
			database.Password,
			database.Host,
			database.Name))

	if err != nil {
		log.Println(err)
	}

	db.SingularTable(true)
	fmt.Println("db ok")
	//	defer db.Close()
}
