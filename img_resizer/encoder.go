package img_resizer

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func JPGEncode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, &jpeg.Options{Quality: jpeg.DefaultQuality})
}

func PNGEncode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

func GIFEncode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}
