package main

import (
	"fmt"
	"lgtm-kinako-api/db"
	"lgtm-kinako-api/domain"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&domain.User{}, &domain.Album{})
}