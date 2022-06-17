package core

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func JpegEncode(source string) (image.Image, error) {
	file, err := os.Open(source)
	if err != nil {
		fmt.Printf("Error opening JPG file %s, %v", source, err)
		return nil, err
	}
	defer file.Close()

	return jpeg.Decode(file)
}

func JpegDecode(dest string, image image.Image) error {
	out, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error creating JPG file %s, %v", dest, err)
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, image, &jpeg.Options{Quality: jpeg.DefaultQuality})
}
