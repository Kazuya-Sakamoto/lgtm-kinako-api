package controller

import (
	"fmt"
	"net/http"

	"lgtm-kinako-api/usecase/album_tag"

	"github.com/labstack/echo/v4"
)

type IAlbumTagController interface {
	ResetAndSetAlbumTags(c echo.Context) error
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

func (ac *albumtagController) ResetAndSetAlbumTags(c echo.Context) error {
	req := new(albumTagRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Error binding request: %s", err.Error()))
	}

	err := ac.atu.ResetAndSetAlbumTags(req.AlbumId, req.TagIds)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error upserting album tags: %s", err.Error()))
	}

	return c.NoContent(http.StatusOK)
}
