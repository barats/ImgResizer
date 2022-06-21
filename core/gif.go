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
