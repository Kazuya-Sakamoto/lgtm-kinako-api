package controller

import (
	"fmt"
	"net/http"

	"lgtm-kinako-api/usecase/album_tag"

	"github.com/labstack/echo/v4"
)

type IAlbumTagController interface {
	DeleteAndInsertAlbumTags(ctx echo.Context) error
	GetAlbumCountsByTag(ctx echo.Context) error
}

type albumtagController struct {
	uc album_tag.IAlbumTagUsecase
}

type albumTagRequest struct {
	AlbumId uint   `json:"albumId"`
	TagIds  []uint `json:"tagIds"`
}

func NewAlbumTagController(uc album_tag.IAlbumTagUsecase) IAlbumTagController {
	return &albumtagController{uc}
}

func (c *albumtagController) DeleteAndInsertAlbumTags(ctx echo.Context) error {
	req := new(albumTagRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %s", err.Error()))
	}

	err := c.uc.DeleteAndInsertAlbumTags(req.AlbumId, req.TagIds)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Error upserting album tags: %s", err.Error()))
	}

	return ctx.NoContent(http.StatusOK)
}

func (c *albumtagController) GetAlbumCountsByTag(ctx echo.Context) error {
	counts, err := c.uc.GetAlbumCountsByTag()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Error retrieving album counts by tag: %s", err.Error()))
	}

	return ctx.JSON(http.StatusOK, counts)
}
