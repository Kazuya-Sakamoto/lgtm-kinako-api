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
	GetTags(ctx echo.Context) error
	CreateTag(ctx echo.Context) error
	DeleteTag(ctx echo.Context) error
}

type tagController struct {
	uc   tag.ITagUsecase
	atuc album_tag.IAlbumTagUsecase
}

func NewTagController(uc tag.ITagUsecase, atuc album_tag.IAlbumTagUsecase) ITagController {
	return &tagController{uc, atuc}
}

func (c *tagController) GetTags(ctx echo.Context) error {
	res, err := c.uc.GetTags()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *tagController) CreateTag(ctx echo.Context) error {
	tag := domain.Tag{}
	if err := ctx.Bind(&tag); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := c.uc.CreateTag(tag)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, res)
}

func (c *tagController) DeleteTag(ctx echo.Context) error {
	id := ctx.Param("tagId")
	tagId, err := strconv.Atoi(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "tagIdが存在しません")
	}

	err = c.atuc.DeleteAlbumTagsByTagID(uint(tagId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	err = c.uc.DeleteTag(uint(tagId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}
