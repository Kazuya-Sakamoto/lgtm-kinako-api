package image_processor

import "image"

type IImageProcessorUsecase interface {
	ProcessImage(inputImage image.Image) ([]byte, error)
	DetectMimeType(data string) (string, error)
}

type ImageProcessorUsecase struct {
	IImageProcessorUsecase
}

func NewImageProcessorUsecase() *ImageProcessorUsecase {
	return &ImageProcessorUsecase{
		NewImageProcessor(),
	}
}
