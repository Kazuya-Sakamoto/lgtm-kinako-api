package image_processor

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/fogleman/gg"
)

const (
	defaultFontSize = 4.8
	defaultText     = "LGTM-kinako"
	fontName        = "Vibur-Regular.ttf"
	targetHeight    = 15.0
)

type imageProcessorUsecase struct{}

func NewImageProcessor() IImageProcessorUsecase {
	return &imageProcessorUsecase{}
}

func (ip *imageProcessorUsecase) ProcessImage(inputImage image.Image) ([]byte, error) {
	dc := gg.NewContextForImage(inputImage)
	dc.SetColor(color.White)

	fontSize := targetHeight * defaultFontSize
	if err := dc.LoadFontFace(fontName, fontSize); err != nil {
		return nil, fmt.Errorf("failed to load font: %v", err)
	}

	if err := ip.drawText(dc); err != nil {
		return nil, fmt.Errorf("failed to draw text: %v", err)
	}

	var buffer bytes.Buffer
	if err := png.Encode(&buffer, dc.Image()); err != nil {
		return nil, fmt.Errorf("failed to encode image: %v", err)
	}

	return buffer.Bytes(), nil
}

func (ip *imageProcessorUsecase) drawText(dc *gg.Context) error {
	text := defaultText
	textWidth, textHeight := dc.MeasureString(text)
	x := (float64(dc.Width()) - textWidth) / 2
	y := (float64(dc.Height()) - textHeight) / 2

	dc.DrawStringAnchored(text, x, y, 0, 0.5)
	return nil
}
