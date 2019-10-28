package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ConnectDB expect
func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "done:done@(localhost)/DONE_development?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	fmt.Println("db connected: ", &db)
	db.LogMode(true)
	return db
}
