package album

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UploadImageToS3Usecase struct{}

func NewUploadImageToS3Usecase() *UploadImageToS3Usecase {
	return &UploadImageToS3Usecase{}
}

func (u *UploadImageToS3Usecase) UploadImageToS3(encodedImage []byte) (string, error) {
	// const s3BaseURL = "https://lgtm-kinako.s3.ap-northeast-1.amazonaws.com/"
	const cloudFrontURL = "https://d18g0hf2wnz3gs.cloudfront.net/"
	awsBucket := os.Getenv("AWS_BUCKET")
	awsProfile := os.Getenv("AWS_PROFILE")
	if awsBucket == "" {
		return "", fmt.Errorf("AWS_BUCKET environment variable is not set")
	}
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           awsProfile,
	})
	if err != nil {
		return "failed to initialize AWS session: %w", err
	}
	svc := s3.New(sess)
	currentDateTime := time.Now().Format("20060102150405")
	imageFileName := currentDateTime + ".JPG"

	// S3に画像をアップロード
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(awsBucket),
		Key:         aws.String(imageFileName),
		Body:        bytes.NewReader(encodedImage),
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload image to S3: %w", err)
	}
	fmt.Println("S3にアップロードが完了")

	objectURL := cloudFrontURL + imageFileName
	return objectURL, nil
}
