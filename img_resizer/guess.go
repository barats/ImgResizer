package img_resizer

import (
	"image"
	"io"
)

func GuessImageType(r io.Reader) (string, error) {
	_, format, err := image.DecodeConfig(r)
	if err != nil {
		return "", err
	}

	return format, nil
}
