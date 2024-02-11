package controller

import (
	"lgtm-kinako-api/usecase/tag"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITagController interface {
	GetTags(c echo.Context) error
}

type tagController struct {
	tu  tag.ITagUsecase
}

func NewTagController(tu tag.ITagUsecase) ITagController {
	return &tagController{tu}
}

func (tc *tagController) GetTags(c echo.Context) error {
	res, err := tc.tu.GetTags()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
