# Lulu
🧙  Lulu(露露) 是一个文件压缩处理的工具箱。
> 使用前，务必先在测试服务器尝试安装 libvips 。

## 依赖
- [libvips](https://www.libvips.org/install.html) 
> 这是一个 C 的动态组件库。如果你的项目是 C，建议不要用这个工具，可能会让一些基础版本不兼容。

## 支持
- 图片压缩

```
$lulu image -h

压缩指定文件夹下的图片，支持多级文件夹。 例如: lulu image

Usage:
  lulu image [flags]

Flags:
  -d, --dir-path string   指定存放图片的文件夹，支持在多级文件夹下获取图片。
  -q, --quality string    指定压缩质量，压缩质量范围为 1(最低质量) - 100(最高质量)，质量越低，图片大小越小。
  -s, --suffix string     指定待处理图片的后缀，目前支持 png，jpeg，jpg。
```

![效果](https://s3.bmp.ovh/imgs/2021/11/de561b2905c114c4.jpg)

## 常见问题

- libvips

  ```
  # pkg-config --cflags vips vips vips vips
  Package vips was not found in the pkg-config search path.
  Perhaps you should add the directory containing `vips.pc'
  to the PKG_CONFIG_PATH environment variable
  No package 'vips' found
  ```
  解决方案：https://www.cxybb.com/article/a_c12345/101725160
