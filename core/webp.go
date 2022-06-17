package core

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"golang.org/x/image/webp"
)

func WebpEncode(source string) (image.Image, error) {
	file, err := os.Open(source)
	if err != nil {
		fmt.Printf("Error opening WEBP file %s, %v", source, err)
		return nil, err
	}
	defer file.Close()

	return webp.Decode(file)
}

func WebpDecode(dest string, image image.Image) error {
	out, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error creating PNG(Converted from WEBP) file %s, %v", dest, err)
		return err
	}
	defer out.Close()

	return png.Encode(out, image)
}
