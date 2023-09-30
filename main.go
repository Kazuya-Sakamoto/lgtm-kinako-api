package main

import (
	"lgtm-kinako-api/controller"
	"lgtm-kinako-api/db"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
	"lgtm-kinako-api/router"
	"lgtm-kinako-api/usecase"
)

func main() {
	db := db.NewDB()
	//* Handler
	userHandler := handler.NewUserHandler()
	albumHandler := handler.NewAlbumHandler()
	// * Repository
	albumRepository := repository.NewAlbumRepository(db)
	userRepository := repository.NewUserRepository(db)
	// * Usecase
	albumUsecase := usecase.NewAlbumUsecase(albumRepository, albumHandler)
	userUsecase := usecase.NewUserUsecase(userRepository, userHandler)
	imageProcessorUsecase := usecase.NewImageProcessor()
	// * Controller
	albumController := controller.NewAlbumController(albumUsecase, imageProcessorUsecase)
	userController := controller.NewUserController(userUsecase)

	e := router.NewRouter(albumController, userController)
	e.Logger.Fatal(e.Start(":8080"))
}