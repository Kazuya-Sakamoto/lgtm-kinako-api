package usecase

import (
	"bytes"
	"fmt"
	"lgtm-kinako-api/handler"
	"lgtm-kinako-api/model"
	"lgtm-kinako-api/repository"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type IAlbumUsecase interface {
	GetAllAlbums() ([]model.AlbumResponse, error)
	GetRandomAlbums() ([]model.AlbumResponse, error)
	CreateAlbum(album model.Album) (model.AlbumResponse, error)
	UploadImageToS3(encodedImage []byte) (string, error)
	DeleteAlbum(userId uint, albumId uint) error
}

type albumUsecase struct {
	ar repository.IAlbumRepository
	ah handler.IAlbumHandler
}

func NewAlbumUsecase(ar repository.IAlbumRepository, ah handler.IAlbumHandler) IAlbumUsecase {
	return &albumUsecase{ar, ah}
}

func (au *albumUsecase) GetAllAlbums() ([]model.AlbumResponse, error) {
	albums := []model.Album{}
	if err := au.ar.GetAllAlbums(&albums); err != nil {
		return nil, err
	}
	res := []model.AlbumResponse{}
	for _, v := range albums {
		a := model.AlbumResponse{
			ID:          v.ID,
			Title:       v.Title,
			Image: 		   v.Image,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		res = append(res, a)
	}
	return res, nil
}

func (au *albumUsecase) GetRandomAlbums() ([]model.AlbumResponse, error) {
	albums := []model.Album{}
	if err := au.ar.GetRandomAlbums(&albums); err != nil {
		return nil, err
	}
	res := []model.AlbumResponse{}
	for _, v := range albums {
		a := model.AlbumResponse{
			ID:          v.ID,
			Title:       v.Title,
			Image: 		   v.Image,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
		res = append(res, a)
	}
	return res, nil
}

func (au *albumUsecase) CreateAlbum(album model.Album) (model.AlbumResponse, error) {
	if err := au.ah.AlbumHandler(album); err != nil {
		return model.AlbumResponse{}, err
	}
	if err := au.ar.CreateAlbum(&album); err != nil {
		return model.AlbumResponse{}, err
	}
	res := model.AlbumResponse{
		ID:          album.ID,
		Title:       album.Title,
		Image: 		 album.Image,
		CreatedAt:   album.CreatedAt,
		UpdatedAt:   album.UpdatedAt,
	}

	return res, nil
}
func (au *albumUsecase) UploadImageToS3(encodedImage []byte) (string, error) {
	const s3BaseURL = "https://lgtm-kinako.s3.ap-northeast-1.amazonaws.com/"
	awsBucket := os.Getenv("AWS_BUCKET")
	awsProfile := os.Getenv("AWS_PROFILE")

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           awsProfile,
	})
	if err != nil {
		return "", err
	}
	svc := s3.New(sess)
	currentDateTime := time.Now().Format("20060102150405")
	imageFileName := currentDateTime + ".JPG"
	//* S3に画像をアップロード
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(awsBucket),
		Key:         aws.String(imageFileName),
		Body:        bytes.NewReader(encodedImage),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return "", err
	}
	fmt.Println("S3にアップロードが完了")

	objectURL := s3BaseURL + imageFileName
	return objectURL, nil
}

func (au *albumUsecase) DeleteAlbum(userId uint, albumId uint) error {
	if err := au.ar.DeleteAlbum(userId, albumId); err != nil {
		return err
	}
	return nil
}
