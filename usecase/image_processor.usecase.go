package usecase

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"strings"

	"github.com/fogleman/gg"
)

type IImageProcessorUsecase interface {
	ProcessImage(inputImage image.Image, targetHeight float64) ([]byte, error)
	DetectMimeType(data string) (string, error)
}

type imageProcessorUsecase struct{}

func NewImageProcessor() IImageProcessorUsecase {
	return &imageProcessorUsecase{}
}

func (ip *imageProcessorUsecase) ProcessImage(inputImage image.Image, targetHeight float64) ([]byte, error) {
	dc := gg.NewContextForImage(inputImage)
	dc.SetColor(color.White)

	fontSize := targetHeight * 4.8
	if err := dc.LoadFontFace("38LSUDGothic-Bold.ttf", fontSize); err != nil {
		fmt.Println("フォントを読み込めませんでした:", err)
		return nil, err
	}

	x := 20.0
	y := (float64(dc.Height()) - fontSize) / 2
	dc.DrawStringAnchored("LGTM-kinako", x, y, 0, 0.5)

	var buffer bytes.Buffer
	if err := png.Encode(&buffer, dc.Image()); err != nil {
		fmt.Println("error3")
		return nil, err
	}

	return buffer.Bytes(), nil
}

func (ip *imageProcessorUsecase) DetectMimeType(data string) (string, error) {
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
