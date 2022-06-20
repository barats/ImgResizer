package main

import (
	"ImgResizer/core"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

const Version string = "v1.2"

var (
	cmdSource     string
	cmdDest       string
	cmdResizeMode int
	cmdWidth      int
	cmdHeight     int
	cmdHelp       bool
	cmdFormat     string
)

func init() {
	flag.StringVar(&cmdFormat, "format", "", "Ouput format. Supported values: png|jpg|jpeg|bmp|tiff|gif. Omit to keep original format.")
	flag.BoolVar(&cmdHelp, "help", false, "Show help message.")
	flag.IntVar(&cmdWidth, "width", -1, "Destination width. Omit to keep original width")
	flag.IntVar(&cmdHeight, "height", -1, "Destination height. Omit to keep original height")
	flag.StringVar(&cmdSource, "source", "", "Source file or directory.")
	flag.StringVar(&cmdDest, "dest", "", "Destination file or directory.")
	flag.IntVar(&cmdResizeMode, "mode", 0, `0 - (Default) Nearest-neighbor interpolation
1 - Bilinear interpolation
2 - Bicubic interpolation
3 - Mitchell-Netravali interpolation
4 - Lanczos resampling with a=2
5 - Lanczos resampling with a=3`)

	flag.Usage = func() {
		fmt.Printf("Usage of ImgResizer %s \nFor more information, please visit: \nhttps://github.com/barats/ImgResizer \nhttps://gitee.com/barat/imgresizer \n\nImgResizer -source {source} -dest {dest} -mode {mode}\n", Version)
		flag.PrintDefaults()
	}
}

func main() {

	flag.Parse()

	if cmdHelp {
		flag.Usage()
		return
	}

	if strings.EqualFold("", strings.TrimSpace(cmdSource)) || strings.EqualFold("", strings.TrimSpace(cmdDest)) {
		fmt.Println("Missing parameter <-source> or <-dest>. Please -h or -help to show help message.")
		return
	}

	sourceInfo, err := os.Stat(cmdSource)
	if err != nil {
		fmt.Printf("Cant not open %s, error %v", cmdSource, err)
		return
	}

	if sourceInfo.IsDir() {
		//Assume that source & destination are directories which include image files in it
		//Assume that destination direcotry is exist(create if it's not)
		//Assume that destination directory is empty(override if it's not)
		files, err := ioutil.ReadDir(cmdSource)
		if err != nil {
			fmt.Printf("Error reading directory %s, %v", cmdSource, err)
			return
		}

		err = os.MkdirAll(cmdDest, os.ModePerm)
		if err != nil {
			fmt.Printf("Error opening or creating directory %s, %v", cmdDest, err)
			return
		}

		for _, f := range files {
			if strings.EqualFold(f.Name(), ".DS_Store") {
				continue
			}
			err := core.DealWithFile(filepath.Join(cmdSource, f.Name()), core.OutputOptions{
				Format:        core.OutputFormat(cmdFormat),
				Width:         cmdWidth,
				Height:        cmdHeight,
				DestPath:      filepath.Join(cmdDest, strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))),
				Interpolation: resize.InterpolationFunction(cmdResizeMode),
			})

			if err != nil {
				fmt.Println(err)
				continue
			}
		} //end of for
	} else {
		//Assume that source & destination is file
		//Assume that destination file does not exist, override if it's not
		err := core.DealWithFile(cmdSource, core.OutputOptions{
			Format:        core.OutputFormat(cmdFormat),
			Width:         cmdWidth,
			Height:        cmdHeight,
			DestPath:      cmdDest,
			Interpolation: resize.InterpolationFunction(cmdResizeMode),
		})

		if err != nil {
			fmt.Sprintln(err)
			return
		}
	}

	fmt.Println("done.")
} //end of main
