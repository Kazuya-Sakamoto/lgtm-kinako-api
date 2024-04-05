package image_processor

import (
	"errors"
	"strings"
)

func (u *imageProcessorUsecase) DetectMimeType(data string) (string, error) {
	parts := strings.SplitN(data, ";", 2)
	if len(parts) != 2 {
		return "", errors.New("無効なデータフォーマット")
	}
	mimeType := strings.TrimSpace(parts[0])
	if !strings.HasPrefix(mimeType, "data:image/") {
		return "", errors.New("無効な画像フォーマット")
	}
	return mimeType, nil
}
