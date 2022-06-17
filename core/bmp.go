package core

import (
	"fmt"
	"image"
	"os"

	"golang.org/x/image/bmp"
)

func BmpEncode(source string) (image.Image, error) {
	file, err := os.Open(source)
	if err != nil {
		fmt.Printf("Error opening BMP file %s, %v", source, err)
		return nil, err
	}
	defer file.Close()

	return bmp.Decode(file)
}

func BmpDecode(dest string, image image.Image) error {
	out, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error creating BMP file %s, %v", dest, err)
		return err
	}
	defer out.Close()

	return bmp.Encode(out, image)
}
