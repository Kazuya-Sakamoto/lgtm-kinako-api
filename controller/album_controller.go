package controller

import (
	"lgtm-kinako-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IAlbumController interface {
	GetAllAlbums(c echo.Context) error
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