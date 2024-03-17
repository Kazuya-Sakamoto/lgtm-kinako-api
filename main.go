package main

import (
	"lgtm-kinako-api/controller"
	"lgtm-kinako-api/db"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/repository"
	"lgtm-kinako-api/router"
	"lgtm-kinako-api/usecase/album"
	"lgtm-kinako-api/usecase/album_tag"
	"lgtm-kinako-api/usecase/image_processor"
	"lgtm-kinako-api/usecase/tag"
	"lgtm-kinako-api/usecase/user"
)

func main() {
	db := db.NewDB()
	/*
		Handler
	*/
	userHandler := handler.NewUserHandler()
	albumHandler := handler.NewAlbumHandler()
	tagHandler := handler.NewTagHandler()
	/*
		Repository
	*/
	albumRepository := repository.NewAlbumRepository(db)
	userRepository := repository.NewUserRepository(db)
	tagRepository := repository.NewTagRepository(db)
	albumTagRepository := repository.NewAlbumTagRepository(db)
	/*
		Usecase
	*/
	albumUsecase := album.NewAlbumUsecase(albumRepository, albumHandler)
	userUsecase := user.NewUserUsecase(userRepository, userHandler)
	imageProcessorUsecase := image_processor.NewImageProcessorUsecase()
	tagUsecase := tag.NewTagUsecase(tagRepository, tagHandler)
	albumTagUsecase := album_tag.NewAlbumTagUsecase(albumTagRepository)
	/*
		Controller
	*/
	albumController := controller.NewAlbumController(albumUsecase, imageProcessorUsecase)
	userController := controller.NewUserController(userUsecase)
	tagController := controller.NewTagController(tagUsecase, albumTagUsecase)
	albumTagController := controller.NewAlbumTagController(albumTagUsecase)

	e := router.NewRouter(albumController, userController, tagController, albumTagController)
	e.Logger.Fatal(e.Start(":8080"))
}
