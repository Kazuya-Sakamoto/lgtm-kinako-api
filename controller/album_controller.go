package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"net/http"
	"strconv"
	"strings"

	"lgtm-kinako-api/domain"
	"lgtm-kinako-api/usecase/album"
	"lgtm-kinako-api/usecase/image_processor"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IAlbumController interface {
	GetAllAlbums(ctx echo.Context) error
	GetRandomAlbums(ctx echo.Context) error
	GetAlbums(ctx echo.Context) error
	CreateAlbum(ctx echo.Context) error
	DeleteAlbum(ctx echo.Context) error
}

type albumController struct {
	uc   album.IAlbumUsecase
	ipuc image_processor.IImageProcessorUsecase
}

func NewAlbumController(uc album.IAlbumUsecase, ipuc image_processor.IImageProcessorUsecase) IAlbumController {
	return &albumController{uc, ipuc}
}

func (c *albumController) GetAllAlbums(ctx echo.Context) error {
	res, err := c.uc.GetAllAlbums()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *albumController) GetRandomAlbums(ctx echo.Context) error {
	res, err := c.uc.GetRandomAlbums()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, res)
}

func (c *albumController) GetAlbums(ctx echo.Context) error {
	id := ctx.QueryParam("tag")
	if id != "" {
		tagId, err := strconv.Atoi(id)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, "Invalid tagId format")
		}
		res, err := c.uc.GetAlbumsByTag(uint(tagId))
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSON(http.StatusOK, res)
	} else {
		res, err := c.uc.GetRandomAlbums()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, res)
	}
}

func (c *albumController) CreateAlbum(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	album := domain.Album{}
	if err := ctx.Bind(&album); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	album.UserId = uint(userId.(float64))

	imageData := album.Image
	mimeType, err := c.ipuc.DetectMimeType(imageData)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	//* MIMEタイプごとに処理を分岐
	if strings.HasPrefix(mimeType, "data:image/jpeg") {
		//* imageDataから"data:image/jpeg;base64,"の部分を削除
		imageData = strings.TrimPrefix(imageData, "data:image/jpeg;base64,")
	} else if strings.HasPrefix(mimeType, "data:image/png") {
		//* imageDataから"data:image/png;base64,"の部分を削除
		imageData = strings.TrimPrefix(imageData, "data:image/png;base64,")
	} else {
		return ctx.JSON(http.StatusBadRequest, "Unsupported image format: jpegでもpngでもありません")
	}

	//* Base64デコード
	data, err := base64.StdEncoding.DecodeString(imageData)
	if err != nil {
		fmt.Println("Base64デコードエラー:", err)
		return ctx.JSON(http.StatusBadRequest, "Unsupported image format: Base64デコードエラー")
	}

	//* デコードしたデータをバイトのストリームとして読み込む
	reader := bytes.NewReader(data)

	decodedImage, format, err := image.Decode(reader)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Unsupported image format: 画像をデコードできません")
	}
	if format != "jpeg" {
		return ctx.JSON(http.StatusBadRequest, "Unsupported image format: フォーマットがJPEGではありません")
	}

	//* 画像の加工
	encodedImage, err := c.ipuc.ProcessImage(decodedImage)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	//* 画像をS3にアップロードし、オブジェクトURLを取得
	objectURL, err := c.uc.UploadImageToS3(encodedImage)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	album.Image = objectURL
	res, err := c.uc.CreateAlbum(album)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, res)
}

func (c *albumController) DeleteAlbum(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := ctx.Param("albumId")
	albumId, _ := strconv.Atoi(id)

	err := c.uc.DeleteAlbum(uint(userId.(float64)), uint(albumId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
