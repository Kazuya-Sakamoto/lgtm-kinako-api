package main

import (
	"fmt"
	"lgtm-kinako-api/db"
	"lgtm-kinako-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Album{})
}