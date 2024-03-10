package album

import (
	"errors"
	"testing"

	"lgtm-kinako-api/repository/mock"

	"github.com/aws/aws-sdk-go/service/s3"
	testify_mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type S3API interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

type MockS3Service struct {
	testify_mock.Mock
}

func (m *MockS3Service) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*s3.PutObjectOutput), args.Error(1)
}

func Test_AlbumUsecase_CreateAlbume_UploadImageToS3(t *testing.T) {
	mockS3 := new(MockS3Service)
	mr := new(mock.MockAlbumRepository)
	mh := new(mock.MockAlbumHandler)
	usecase := NewUploadImageToS3Usecase(mr, mh)

	t.Run("画像が正常にアップロードされること", func(t *testing.T) {
		/*
			正常系のテストを書く
		*/
	})

	t.Run("S3へのアップロードでエラーが発生する場合", func(t *testing.T) {
		encodedImage := []byte("test image data")

		mockS3.On("PutObject", testify_mock.Anything).Return(nil, errors.New("s3 error"))

		_, err := usecase.UploadImageToS3(encodedImage)

		require.Error(t, err)
	})
}
