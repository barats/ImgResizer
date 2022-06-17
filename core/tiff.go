package core

import (
	"fmt"
	"image"
	"os"

	"golang.org/x/image/tiff"
)

func TiffEncode(source string) (image.Image, error) {
	file, err := os.Open(source)
	if err != nil {
		fmt.Printf("Error opening TIFF file %s, %v", source, err)
		return nil, err
	}
	defer file.Close()

	return tiff.Decode(file)
}

func TiffDecode(dest string, image image.Image) error {
	out, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error creating TIFF file %s, %v", dest, err)
		return err
	}
	defer out.Close()

	return tiff.Encode(out, image, &tiff.Options{Compression: tiff.Uncompressed})
}
