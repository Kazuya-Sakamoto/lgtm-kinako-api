package controller

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/usecase"
	"net/http"
	"strconv"
	"strings"

	"github.com/fogleman/gg"

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
	au usecase.IAlbumUsecase
}

func NewAlbumController(au usecase.IAlbumUsecase) IAlbumController {
	return &albumController{au}
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
	mimeType, err := detectMimeType(imageData)
	if err != nil {
		fmt.Println("error100")
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// MIMEタイプごとに処理を分岐
	if strings.HasPrefix(mimeType, "data:image/jpeg") {
		// JPEG形式の画像データの処理
		// imageDataから"data:image/jpeg;base64,"の部分を削除
		imageData = strings.TrimPrefix(imageData, "data:image/jpeg;base64,")
	} else if strings.HasPrefix(mimeType, "data:image/png") {
		// PNG形式の画像データの処理
		// imageDataから"data:image/png;base64,"の部分を削除
		imageData = strings.TrimPrefix(imageData, "data:image/png;base64,")
	} else {
		return c.JSON(http.StatusBadRequest, "Unsupported image format")
	}

    // Base64デコード
    data, err := base64.StdEncoding.DecodeString(imageData)
    if err != nil {
        fmt.Println("Base64デコードエラー:", err)
        return c.JSON(http.StatusBadRequest, "Unsupported image format")
    }

    // デコードしたデータをバイトのストリームとして読み込む
    reader := bytes.NewReader(data)

    // JPEG形式の画像としてデコード
    decodedImage, format, err := image.Decode(reader)
    if err != nil {
        fmt.Println("画像デコードエラー:", err)
        return c.JSON(http.StatusBadRequest, "Unsupported image format")
    }

    // フォーマットがJPEGであるか確認
    if format != "jpeg" {
        fmt.Println("無効な画像フォーマット:", format)
        return c.JSON(http.StatusBadRequest, "Unsupported image format")
    }

    fmt.Println("画像が正常にデコードされました。")

    // 画像を加工してBase64エンコード
	targetHeight := 15.0 // 15センチの高さを目指す
    encodedImage, err := processImage(decodedImage, targetHeight)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    album.Image = encodedImage

    res, err := ac.au.CreateAlbum(album)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }

    return c.JSON(http.StatusCreated, res)
}

// detectMimeType 関数はここで定義
func detectMimeType(data string) (string, error) {
	parts := strings.SplitN(data, ";", 2)
	if len(parts) != 2 {
			return "", errors.New("Invalid data format")
	}
	mimeType := strings.TrimSpace(parts[0])
	if !strings.HasPrefix(mimeType, "data:image/") {
			return "", errors.New("Invalid image format")
	}
	return mimeType, nil
}

func processImage(inputImage image.Image, targetHeight float64) (string, error) {
	// 画像にテキストを追加
	dc := gg.NewContextForImage(inputImage)
	dc.SetColor(color.White)

	// テキストのフォントサイズを5倍に設定
	fontSize := targetHeight * 4.8
	if err := dc.LoadFontFace("38LSUDGothic-Bold.ttf", fontSize); err != nil {
			fmt.Println("フォントを読み込めませんでした:", err)
			return "", err
	}

	// テキストを左上の固定位置に配置
	x := 20.0 // 画像の左端からの距離（20px）
	y := (float64(dc.Height()) - fontSize) / 2

	// テキストを配置
	dc.DrawStringAnchored("LGTM-kinako", x, y, 0, 0.5)

	// 加工した画像をバッファに書き込む
	var buffer bytes.Buffer
	if err := png.Encode(&buffer, dc.Image()); err != nil {
			fmt.Println("error3")
			return "", err
	}

	// バッファからBase64エンコードされた文字列に変換
	encodedImage := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return encodedImage, nil
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
