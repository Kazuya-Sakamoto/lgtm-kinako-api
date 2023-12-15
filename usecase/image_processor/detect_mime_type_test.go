package image_processor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetectMimeType(t *testing.T) {
	ip := &imageProcessorUsecase{}

	t.Run("有効な画像データの場合", func(t *testing.T) {
		data := "data:image/jpeg;base64,iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg=="
		mimeType, err := ip.DetectMimeType(data)
		require.NoError(t, err)
		require.Equal(t, "data:image/jpeg", mimeType)
	})	

	t.Run("無効なデータフォーマットの場合", func(t *testing.T) {
		data := "invalid_data_format"
		_, err := ip.DetectMimeType(data)
		require.Error(t, err)
		require.Equal(t, "無効なデータフォーマット", err.Error())
	})

	t.Run("無効な画像フォーマットの場合", func(t *testing.T) {
		data := "data:text/plain;base64,SGVsbG8sIFdvcmxkIQ=="
		_, err := ip.DetectMimeType(data)
		require.Error(t, err)
		require.Equal(t, "無効な画像フォーマット", err.Error())
	})
}
