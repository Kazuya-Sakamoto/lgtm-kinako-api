package controller

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


func processImage(inputImage image.Image, targetHeight float64) ([]byte, error) {
	// 画像にテキストを追加
	dc := gg.NewContextForImage(inputImage)
	dc.SetColor(color.White)

	// テキストのフォントサイズを5倍に設定
	fontSize := targetHeight * 4.8
	if err := dc.LoadFontFace("38LSUDGothic-Bold.ttf", fontSize); err != nil {
		fmt.Println("フォントを読み込めませんでした:", err)
		return nil, err
	}

	// テキストを左上の固定位置に配置
	x := 20.0 // 画像の左端からの距離（20px）
	y := (float64(dc.Height()) - fontSize) / 2

	// テキストを配置
	dc.DrawStringAnchored("LGTM-kinako", x, y, 0, 0.5)

	// 加工した画像をバッファに書き込む
	var buffer bytes.Buffer
	if err := png.Encode(&buffer, dc.Image()); err != nil {
		fmt.Println("error3")
		return nil, err
	}

	// バッファの内容を []byte として返す
	return buffer.Bytes(), nil
}

func detectMimeType(data string) (string, error) {
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
