package core

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func PngEncode(source string) (image.Image, error) {
	file, err := os.Open(source)
	if err != nil {
		fmt.Printf("Error opening PNG file %s, %v", source, err)
		return nil, err
	}
	defer file.Close()

	return png.Decode(file)
}

func PngDecode(dest string, image image.Image) error {
	out, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error creating PNG file %s, %v", dest, err)
		return err
	}
	defer out.Close()

	return png.Encode(out, image)
}
