//Copyright (c) [2022] [巴拉迪维]
//[ImgResizer] is licensed under Mulan PSL v2.
//You can use this software according to the terms and conditions of the Mulan PSL v2.
//You may obtain a copy of Mulan PSL v2 at:
//http://license.coscl.org.cn/MulanPSL2
//THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
//See the Mulan PSL v2 for more details.

package core

import (
	"fmt"
	"image"
	"io"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

type OutputOptions struct {
	Width         int
	Height        int
	Interpolation resize.InterpolationFunction
	DestPath      string
	Format        OutputFormat
}

type OutputFormat string

const (
	PNG  OutputFormat = "png"
	JPG  OutputFormat = "jpg"
	JPEG OutputFormat = "jpeg"
	BMP  OutputFormat = "bmp"
	TIFF OutputFormat = "tiff"
	GIF  OutputFormat = "gif"
)

//
//Deal with image files
func DealWithFile(source string, option OutputOptions) error {
	file, err := os.Open(source)
	if err != nil {
		fmt.Printf("Error opening file %s, %v", source, err)
		return err
	}
	defer file.Close()

	fileFormat, width, height, err := retrieveImageInfo(file)
	if err != nil {
		fmt.Printf("Could not guess mime type of file %s, %v", source, err)
		return err
	}

	file.Seek(0, 0) //MUST SEEK BACK TO 0,0 acording to https://github.com/golang/go/issues/50992

	var data image.Image
	//encode image file for different type
	switch fileFormat {
	case "bmp":
		data, err = BmpEncode(source)
	case "jpg":
		data, err = JpegEncode(source)
	case "jpeg":
		data, err = JpegEncode(source)
	case "tiff":
		data, err = TiffEncode(source)
	case "webp":
		data, err = WebpEncode(source)
	case "png":
		data, err = PngEncode(source)
	case "gif":
		data, err = GifEncode(source)
	default:
		err = fmt.Errorf("unsupported image type %s", fileFormat)
	}

	if err != nil {
		return err
	}

	if option.Width == -1 {
		option.Width = width
	}

	if option.Height == -1 {
		option.Height = height
	}

	afterResize := resize.Resize(uint(option.Width), uint(option.Height), data, option.Interpolation)

	if strings.EqualFold("", string(option.Format)) {
		option.Format = OutputFormat(fileFormat)
	}

	switch option.Format {
	case "bmp":
		err = BmpDecode(fmt.Sprintf("%s_%d_%d.bmp", strings.TrimSpace(option.DestPath), option.Width, option.Height), afterResize)
	case "jpg":
		err = JpegDecode(fmt.Sprintf("%s_%d_%d.jpg", strings.TrimSpace(option.DestPath), option.Width, option.Height), afterResize)
	case "jpeg":
		err = JpegDecode(fmt.Sprintf("%s_%d_%d.jpeg", strings.TrimSpace(option.DestPath), option.Width, option.Height), afterResize)
	case "tiff":
		err = TiffDecode(fmt.Sprintf("%s_%d_%d.tiff", strings.TrimSpace(option.DestPath), option.Width, option.Height), afterResize)
	case "gif":
		err = GifDecode(fmt.Sprintf("%s_%d_%d.gif", strings.TrimSpace(option.DestPath), option.Width, option.Height), afterResize)
	case "webp":
	case "png":
		err = PngDecode(fmt.Sprintf("%s_%d_%d.png", strings.TrimSpace(option.DestPath), option.Width, option.Height), afterResize)
	default:
		err = fmt.Errorf("unsupported output format %s", option.Format)
	}

	return err
} //end of function

//
//Guess image type
func retrieveImageInfo(r io.Reader) (string, int, int, error) {
	config, format, err := image.DecodeConfig(r)
	if err != nil {
		return "", -1, -1, err
	}

	return format, config.Width, config.Height, nil
}
