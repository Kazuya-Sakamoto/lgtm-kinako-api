package main

import (
	"lgtm-kinako-api/controller"
	"lgtm-kinako-api/db"
	"lgtm-kinako-api/repository"
	"lgtm-kinako-api/router"
	"lgtm-kinako-api/usecase"
)

func main() {
	db := db.NewDB()
	alubmRepository := repository.NewAlbumRepository(db)
	alubmUsecase := usecase.NewAlbumUsecase(alubmRepository)
	alubmController := controller.NewAlbumController(alubmUsecase)

	e := router.NewRouter(alubmController)
	e.Logger.Fatal(e.Start(":8080"))
}