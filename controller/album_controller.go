package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/usecase"
	"lgtm-kinako-api/usecase/image_processor"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IAlbumController interface {
	GetAllAlbums(c echo.Context) error
	GetRandomAlbums(c echo.Context) error
	CreateAlbum(c echo.Context) error
	DeleteAlbum(c echo.Context) error
}

type albumController struct {
	au  usecase.IAlbumUsecase
	ipu image_processor.IImageProcessorUsecase
}

func NewAlbumController(au usecase.IAlbumUsecase, ipu image_processor.IImageProcessorUsecase) IAlbumController {
	return &albumController{au, ipu}
}

func (ac *albumController) GetAllAlbums(c echo.Context) error {
	res, err := ac.au.GetAllAlbums()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (ac *albumController) GetRandomAlbums(c echo.Context) error {
	res, err := ac.au.GetRandomAlbums()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}


func (ac *albumController) CreateAlbum(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	album := model.Album{}
	if err := c.Bind(&album); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	album.UserId = uint(userId.(float64))

	imageData := album.Image
	mimeType, err := ac.ipu.DetectMimeType(imageData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//* MIMEタイプごとに処理を分岐
	if strings.HasPrefix(mimeType, "data:image/jpeg") {
		//* imageDataから"data:image/jpeg;base64,"の部分を削除
		imageData = strings.TrimPrefix(imageData, "data:image/jpeg;base64,")
	} else if strings.HasPrefix(mimeType, "data:image/png") {
		//* imageDataから"data:image/png;base64,"の部分を削除
		imageData = strings.TrimPrefix(imageData, "data:image/png;base64,")
	} else {
		return c.JSON(http.StatusBadRequest, "Unsupported image format: jpegでもpngでもありません")
	}

	//* Base64デコード
	data, err := base64.StdEncoding.DecodeString(imageData)
	if err != nil {
		fmt.Println("Base64デコードエラー:", err)
		return c.JSON(http.StatusBadRequest, "Unsupported image format: Base64デコードエラー")
	}

	//* デコードしたデータをバイトのストリームとして読み込む
	reader := bytes.NewReader(data)

	decodedImage, format, err := image.Decode(reader)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Unsupported image format: 画像をデコードできません")
	}
	if format != "jpeg" {
		return c.JSON(http.StatusBadRequest, "Unsupported image format: フォーマットがJPEGではありません")
	}

	//* アップロードする画像データ
	encodedImage, err := ac.ipu.ProcessImage(decodedImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//* 画像をS3にアップロードし、オブジェクトURLを取得
	objectURL, err := ac.au.UploadImageToS3(encodedImage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
    album.Image = objectURL
    res, err := ac.au.CreateAlbum(album)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, res)
}


func (ac *albumController) DeleteAlbum(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("albumId")
	albumId, _ := strconv.Atoi(id)

	err := ac.au.DeleteAlbum(uint(userId.(float64)), uint(albumId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
