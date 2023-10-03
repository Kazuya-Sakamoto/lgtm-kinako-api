package main

import (
	"lgtm-kinako-api/controller"
	"lgtm-kinako-api/db"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
	"lgtm-kinako-api/router"
	"lgtm-kinako-api/usecase"
	"lgtm-kinako-api/usecase/image_processor"
	"lgtm-kinako-api/usecase/user"
)

func main() {
	db := db.NewDB()
	// * Handler
	userHandler := handler.NewUserHandler()
	albumHandler := handler.NewAlbumHandler()
	// * Repository
	albumRepository := repository.NewAlbumRepository(db)
	userRepository := repository.NewUserRepository(db)
	// * Usecase
	albumUsecase := usecase.NewAlbumUsecase(albumRepository, albumHandler)
	userUsecase := user.NewUserUsecase(userRepository, userHandler)
	imageProcessorUsecase := image_processor.NewImageProcessorUsecase()
	// * Controller
	albumController := controller.NewAlbumController(albumUsecase, imageProcessorUsecase)
	userController := controller.NewUserController(userUsecase)

	e := router.NewRouter(albumController, userController)
	e.Logger.Fatal(e.Start(":8080"))
}