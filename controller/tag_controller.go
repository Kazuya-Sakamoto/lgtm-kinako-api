package controller

import (
	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/usecase/tag"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITagController interface {
	GetTags(c echo.Context) error
	CreateTag(c echo.Context) error
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

func (tc *tagController) CreateTag(c echo.Context) error {
	tag := domain.Tag{}
	if err := c.Bind(&tag); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := tc.tu.CreateTag(tag)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, res)
}
