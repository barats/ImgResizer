package img_resizer

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func JPGDecode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

func PNGDecode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func GIFDecode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}
