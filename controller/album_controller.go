package controller

import (
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IAlbumController interface {
	GetAllAlbums(c echo.Context) error
	CreateAlbum(c echo.Context) error
	DeleteAlbum(c echo.Context) error
}

type albumController struct {
	au usecase.IAlbumUsecase
}

func NewAlbumController(au usecase.IAlbumUsecase) IAlbumController {
	return &albumController{au}
}

func (ac *albumController) GetAllAlbums(c echo.Context) error {
	res, err := ac.au.GetAllAlbums()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (ac *albumController) CreateAlbum(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	album := model.Album{}
	if err := c.Bind(&album); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	album.UserId = uint(userId.(float64))
	res, err := ac.au.CreateAlbum(album)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, res)
}

func (ac *albumController) DeleteAlbum(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("albumId")
	albumId, _ := strconv.Atoi(id)

	err := ac.au.DeleteAlbum(uint(userId.(float64)), uint(albumId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
