package controller

import (
	"net/http"
	"strconv"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/usecase/album_tag"
	"lgtm-kinako-api/usecase/tag"

	"github.com/labstack/echo/v4"
)

type ITagController interface {
	GetTags(c echo.Context) error
	CreateTag(c echo.Context) error
	DeleteTag(c echo.Context) error
}

type tagController struct {
	tu  tag.ITagUsecase
	atu album_tag.IAlbumTagUsecase
}

func NewTagController(tu tag.ITagUsecase, atu album_tag.IAlbumTagUsecase) ITagController {
	return &tagController{tu, atu}
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

func (tc *tagController) DeleteTag(c echo.Context) error {
	id := c.Param("tagId")
	tagId, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "tagIdが存在しません")
	}

	err = tc.atu.DeleteAlbumTagsByTagId(uint(tagId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = tc.tu.DeleteTag(uint(tagId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
