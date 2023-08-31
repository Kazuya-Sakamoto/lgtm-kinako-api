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
	userRepository := repository.NewUserRepository(db)
	alubmUsecase := usecase.NewAlbumUsecase(alubmRepository)
	userUsecase := usecase.NewUserUsecase(userRepository)
	alubmController := controller.NewAlbumController(alubmUsecase)
	userController := controller.NewUserController(userUsecase)

	e := router.NewRouter(alubmController, userController)
	e.Logger.Fatal(e.Start(":8080"))
}