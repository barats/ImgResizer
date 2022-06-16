# ImgResizer

批量图片等比缩放工具，同时支持单独文件  
支持 png 格式、gif 格式、jpeg 格式

## 使用方法

```
Usage: ImgResizer -source {source} -dest {dest} -mode {mode}
  -dest string
    	Destination file or directory
  -height int
    	Destination height (default 128)
  -help
    	Show help message
  -mode int
    	0 - (Default) Nearest-neighbor interpolation
    	1 - Bilinear interpolation
    	2 - Bicubic interpolation
    	3 - Mitchell-Netravali interpolation
    	4 - Lanczos resampling with a=2
    	5 - Lanczos resampling with a=3

  -source string
    	Source file or directory
  -width int
    	Destination width (default 300)
```

## 示例批量处理

```
ImgResizer -source ~/Desktop/pics -dest ~/Desktop/new_pics -mode 5
```

## 示例单独处理

```
ImgResizer -source ~/pics/hello.gif -dest ~/newpics/wow.gif -width 900
```

