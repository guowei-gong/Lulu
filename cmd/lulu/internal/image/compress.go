package image

import (
	"fmt"
	"github.com/h2non/bimg"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	FileSep    = "." // 文件名切割符号
	PackageSep = "/" // 文件夹切割符号

	PNG  = "png"
	JPG  = "jpg"
	JPEG = "jpeg"
)

func New(suffix, dirPath, qualityStr string) error {
	quality, err := strconv.Atoi(qualityStr)

	if err != nil {
		return err
	}

	files, _ := WalkDir(dirPath, suffix)

	for _, read := range files {
		fmt.Printf("🍺  %s 开始处理\n", read)

		// 获取图片路径
		write := WriteUrl(read)

		// 格式转换
		if suffix == PNG || suffix == JPG {
			Convert(read, write)
		}
		// 图片压缩
		Compress(write, read, quality)

		// 删除临时生成的图片
		if suffix != JPEG {
			Delete(write)
		}
	}

	fmt.Println("🤝 Thanks for using Compress")
	return nil
}

func Delete(url string) error {
	return os.Remove(url)
}

func WriteUrl(read string) string {
	directory := strings.Split(read, PackageSep)

	old := directory[len(directory)-1]

	fileName := strings.Split(old, FileSep)
	replaceNew := fileName[0] + FileSep + JPEG

	return strings.Replace(read, old, replaceNew, -1)
}

func WalkDir(dirPath, suffix string) (files []string, err error) {
	files = make([]string, 0)
	suffix = strings.ToUpper(suffix)
	err = filepath.Walk(dirPath, func(filename string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(info.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

func Compress(read, write string, quality int) string {
	options := bimg.Options{
		Quality: quality,
	}

	buffer, err := bimg.Read(read)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Process(options)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	err = bimg.Write(write, newImage)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return ""
}

func Convert(read, write string) {
	buffer, err := bimg.Read(read)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Convert(bimg.JPEG)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	//if bimg.NewImage(newImage).Type() == JPEG {
	//	fmt.Fprintln(os.Stderr, "The image was converted into jpeg")
	//}

	bimg.Write(write, newImage)
}
