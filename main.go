package main

import (
	"ImgResize/img_resizer"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

var (
	cmdSource     string
	cmdDest       string
	cmdResizeMode int
	cmdWidth      int
	cmdHeight     int
	cmdHelp       bool
)

func init() {
	flag.BoolVar(&cmdHelp, "help", false, "Show help message")
	flag.IntVar(&cmdWidth, "width", 300, "Destination width")
	flag.IntVar(&cmdHeight, "height", 128, "Destination height")
	flag.StringVar(&cmdSource, "source", "", "Source file or directory")
	flag.StringVar(&cmdDest, "dest", "", "Destination file or directory")
	flag.IntVar(&cmdResizeMode, "mode", 0, `0 - (Default) Nearest-neighbor interpolation
1 - Bilinear interpolation
2 - Bicubic interpolation
3 - Mitchell-Netravali interpolation
4 - Lanczos resampling with a=2
5 - Lanczos resampling with a=3
`)

	flag.Usage = func() {
		fmt.Println("Usage: ImgResizer -source {source} -dest {dest} -mode {mode}")
		flag.CommandLine.PrintDefaults()
	}
}

func main() {

	flag.Parse()

	if cmdHelp {
		flag.Usage()
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
			sf := filepath.Join(cmdSource, f.Name())
			df := filepath.Join(cmdDest, strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())))
			if strings.EqualFold(f.Name(), ".DS_Store") {
				continue
			}
			err := dealWithFile(sf, df)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} //end of for

	} else {
		//Assume that source & destination is file
		//Assume that destination file does not exist, override if it's not
		err := dealWithFile(cmdSource, cmdDest)
		if err != nil {
			return
		}
	}

	fmt.Println("done.")

} //end of main

func dealWithFile(source, dest string) error {

	ofile, err := os.Open(source)
	if err != nil {
		fmt.Printf("Error opening file %s, %v", source, err)
		return err
	}
	defer ofile.Close()

	format, err := img_resizer.GuessImageType(ofile)
	if err != nil {
		fmt.Printf("Could not guess mime type of file %s, %v", source, err)
		return err
	}

	ofile.Seek(0, 0) //MUST SEEK BACK TO 0,0 acording to https://github.com/golang/go/issues/50992

	if strings.EqualFold("png", strings.ToLower(format)) {
		img, err := img_resizer.PNGDecode(ofile)
		if err != nil {
			fmt.Printf("Error decoding PNG file %s, %v", source, err)
			return err
		}

		destFile := resize.Resize(uint(cmdWidth), uint(cmdHeight), img, resize.InterpolationFunction(cmdResizeMode))

		out, err := os.Create(fmt.Sprintf("%s_%d_%d.png", strings.TrimSpace(dest), cmdWidth, cmdHeight))
		if err != nil {
			fmt.Printf("Error creating file %s, %v", dest, err)
			return err
		}
		defer out.Close()

		return img_resizer.PNGEncode(out, destFile)
	}

	if strings.EqualFold("jpeg", strings.ToLower(format)) || strings.EqualFold("jpg", strings.ToLower(format)) {
		img, err := img_resizer.JPGDecode(ofile)
		if err != nil {
			fmt.Printf("Error decoding JPEG file %s, %v", source, err)
			return err
		}

		destFile := resize.Resize(uint(cmdWidth), uint(cmdHeight), img, resize.InterpolationFunction(cmdResizeMode))

		out, err := os.Create(fmt.Sprintf("%s_%d_%d.jpg", strings.TrimSpace(dest), cmdWidth, cmdHeight))
		if err != nil {
			fmt.Printf("Error creating file %s, %v", dest, err)
			return err
		}
		defer out.Close()

		return img_resizer.JPGEncode(out, destFile)
	}

	if strings.EqualFold("gif", strings.ToLower(format)) {
		img, err := img_resizer.GIFDecode(ofile)
		if err != nil {
			fmt.Printf("Error decoding GIF file %s, %v", source, err)
			return err
		}

		destFile := resize.Resize(uint(cmdWidth), uint(cmdHeight), img, resize.InterpolationFunction(cmdResizeMode))

		out, err := os.Create(fmt.Sprintf("%s_%d_%d.gif", strings.TrimSpace(dest), cmdWidth, cmdHeight))
		if err != nil {
			fmt.Printf("Error creating file %s, %v", dest, err)
			return err
		}
		defer out.Close()

		return img_resizer.GIFEncode(out, destFile)
	}

	return nil
}
