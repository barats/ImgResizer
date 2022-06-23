# ImgResizer

批量图片等比缩放、类型转换工具  
1. 支持图片类型：bmp、tiff、jpg、jpeg、gif、png、webp  
1. 支持类型转换为：bmp、tiff、jpg、jpeg、gif、png  
1. 支持自定义宽度、高度
1. 五种等比缩放模式

<p style="text-align: center">
<a target="_blank" href="https://github.com/barats/ImgResizer/stargazers"><img src="https://img.shields.io/github/stars/barats/ImgResizer"/></a>
<a target="_blank" href="https://github.com/barats/ImgResizer/network/members"><img src="https://img.shields.io/github/forks/barats/ImgResizer"/></a>
<a target="_blank" href="https://github.com/barats/ImgResizer/issues"><img src="https://img.shields.io/github/issues/barats/ImgResizer"/></a>  
<a target="_blank" href='https://gitee.com/barat/imgresizer/stargazers'><img src='https://gitee.com/barat/imgresizer/badge/star.svg?theme=dark' /></a>
<a target="_blank" href='https://gitee.com/barat/imgresizer/members'><img src='https://gitee.com/barat/imgresizer/badge/fork.svg?theme=dark' /></a>
<a target="_blank" href='https://www.oschina.net/comment/project/64382'><img src='https://www.oschina.net/comment/badge/project/64382'/></a> 
</p>

## 使用方法

```
ImgResizer -source {source} -dest {dest} -mode {mode}
  -dest string
        Destination file or directory
  -format string
        Output format 
        Supported values: png|jpg|jpeg|bmp|tiff|gif 
        Omit to keep original format 
  -height int
        Destination height 
        Omit to keep original height (default -1)
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
        Destination width tai
        Omit to keep original width (default -1)
```

## 注意事项

1. **如果不需要改变原图类型，请省略 `-format` 参数**
1. **webp 格式图片，默认转换为 png 格式处理(目前没有 webp 图片的高效、简洁处理办法)**
1. **如果不需要改变原图尺寸，请同时省略 `-width` 和 `-height` 参数** 

## 使用示例

### 1. 批量等比缩放

```
ImgResizer -source ~/pics -dest ~/new_pics -mode 5 -height 128 -width 300
```

### 2. 单文件指定宽度

```
ImgResizer -source ~/pics/hello.gif -dest ~/newpics/wow.gif -width 900
```

### 3. 批量类型转换

```
ImgResizer -source ~/pics -dest ~/new_pics -format jpg
```

## ImgResizer 
1. Gitee [https://gitee.com/barat/imgresizer](https://gitee.com/barat/imgresizer)
1. Github [https://github.com/barats/ImgResizer](https://github.com/barats/ImgResizer)
1. Gitlink [https://www.gitlink.org.cn/baladiwei/imgresizer](https://www.gitlink.org.cn/baladiwei/imgresizer)
1. 收录信息  [https://www.oschina.net/p/ImgResizer](https://www.oschina.net/p/ImgResizer)