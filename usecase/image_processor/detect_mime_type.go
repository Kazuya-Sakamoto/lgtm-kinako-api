package image_processor

import (
	"errors"
	"strings"
)

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
