package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/usecase"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

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
	const s3BaseURL = "https://lgtm-kinako.s3.ap-northeast-1.amazonaws.com/"
    // AWSセッションを作成
    // デフォルトのAWS設定ファイル（~/.aws/credentials）から認証情報を読み込む
    // 指定のプロファイルを使用する場合は、第二引数にプロファイル名を指定します
    sess, err := session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
        Profile:           "kinako-lgtm", 
    })
    if err != nil {
        panic(err)
    }

    // S3クライアントを作成
    svc := s3.New(sess)

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


    // イメージファイルの名前を作成（現在の日付と時間を使用）
    currentDateTime := time.Now().Format("20060102150405") // YYYYMMDDHHMMSS 形式
    imageFileName := currentDateTime + ".JPG"

    // アップロードする画像データ
	encodedImage, err := processImage(decodedImage, 15.0)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

    // S3にアップロード
    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket:      aws.String("lgtm-kinako"),           	// バケット名を指定
        Key:         aws.String(imageFileName),           	// 保存するS3キーを指定
        Body:        bytes.NewReader([]byte(encodedImage)), // encodedImageをバイトスライスに変換して指定
        ContentType: aws.String("image/jpeg"),        		// Content-Typeを指定
    })
    if err != nil {
        panic(err)
    }
    fmt.Println("画像をS3に保存しました。")

    // 保存された画像のオブジェクト URL を生成
	objectURL := s3BaseURL + imageFileName
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
