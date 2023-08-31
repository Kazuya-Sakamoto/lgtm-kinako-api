package router

import (
	"lgtm-kinako-api/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(ac controller.IAlbumController) *echo.Echo {
	e := echo.New()

	a := e.Group("/albums")
	a.GET("", ac.GetAllAlbums)
	return e
}