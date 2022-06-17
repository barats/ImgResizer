# ImgResizer

批量图片等比缩放、类型转换工具  
1. 支持图片类型：bmp、tiff、jpg、jpeg、gif、png、webp  
1. 支持类型转换为：bmp、tiff、jpg、jpeg、gif、png  
1. 支持自定义宽度、高度  
1. 五种等比缩放模式

## 使用方法

```
ImgResizer -source {source} -dest {dest} -mode {mode}

  -dest string
        Destination file or directory
  -format string
        Ouput format (png|jpg|jpeg|bmp|tiff|gif)
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

## 批量等比缩放处理

```
ImgResizer -source ~/Desktop/pics -dest ~/Desktop/new_pics -mode 5
```

## 单个文件指定宽度缩放

```
ImgResizer -source ~/pics/hello.gif -dest ~/newpics/wow.gif -width 900
```

## 批量文件类型转换
```
ImgResizer -source ~/Desktop/pics -dest ~/Desktop/new_pics -format jpg
```

