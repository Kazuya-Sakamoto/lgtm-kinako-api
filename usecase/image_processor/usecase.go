package image_processor

import "image"

type ImageProcessorUsecase struct {
	ProcessorUsecase IImageProcessorUsecase
}

func NewImageProcessorUsecase() *ImageProcessorUsecase {
	return &ImageProcessorUsecase{
		ProcessorUsecase: NewImageProcessor(),
	}
}

func (ipu *ImageProcessorUsecase) ProcessImage(inputImage image.Image) ([]byte, error) {
	return ipu.ProcessorUsecase.ProcessImage(inputImage)
}

func (ipu *ImageProcessorUsecase) DetectMimeType(data string) (string, error) {
	return ipu.ProcessorUsecase.DetectMimeType(data)
}
