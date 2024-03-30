package controller

import (
	"fmt"
	"net/http"

	"lgtm-kinako-api/usecase/album_tag"

	"github.com/labstack/echo/v4"
)

type IAlbumTagController interface {
	DeleteAndInsertAlbumTags(c echo.Context) error
	GetAlbumCountsByTag(c echo.Context) error
}

type albumtagController struct {
	atu album_tag.IAlbumTagUsecase
}

type albumTagRequest struct {
	AlbumId uint   `json:"albumId"`
	TagIds  []uint `json:"tagIds"`
}

func NewAlbumTagController(atu album_tag.IAlbumTagUsecase) IAlbumTagController {
	return &albumtagController{atu}
}

func (ac *albumtagController) DeleteAndInsertAlbumTags(c echo.Context) error {
	req := new(albumTagRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %s", err.Error()))
	}

	err := ac.atu.DeleteAndInsertAlbumTags(req.AlbumId, req.TagIds)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error upserting album tags: %s", err.Error()))
	}

	return c.NoContent(http.StatusOK)
}

func (ac *albumtagController) GetAlbumCountsByTag(c echo.Context) error {
	counts, err := ac.atu.GetAlbumCountsByTag()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error retrieving album counts by tag: %s", err.Error()))
	}

	return c.JSON(http.StatusOK, counts)
}
