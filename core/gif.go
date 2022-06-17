package core

import (
	"fmt"
	"image"
	"image/gif"
	"os"
)

func GifEncode(source string) (image.Image, error) {
	file, err := os.Open(source)
	if err != nil {
		fmt.Printf("Error opening GIF file %s, %v", source, err)
		return nil, err
	}
	defer file.Close()

	return gif.Decode(file)
}

func GifDecode(dest string, image image.Image) error {
	out, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error creating GIF file %s, %v", dest, err)
		return err
	}
	defer out.Close()

	return gif.Encode(out, image, nil)
}
