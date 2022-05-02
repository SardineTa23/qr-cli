package qrgen

import (
	"bytes"
	"image"
	"image/png"
)

func CreateImage(b []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(b))
}
