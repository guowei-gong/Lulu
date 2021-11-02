package image

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"os"
)

var CmdImage = &cobra.Command{
	Use:   "image",
	Short: "压缩图片.",
	Long:  "压缩指定文件夹下的图片, 支持多级文件夹. 例如: lulu image",
	Run:   run,
}

var (
	dirPath string
	suffix  string
	quality string
)

func init() {
	CmdImage.Flags().StringVarP(&dirPath, "dir-path", "d", dirPath, "指定存放图片的文件夹, 支持在多级文件夹下获取图片.")
	CmdImage.Flags().StringVarP(&suffix, "suffix", "s", suffix, "指定待处理图片的后缀, 目前支持 png, jpeg, jpg.")
	CmdImage.Flags().StringVarP(&quality, "quality", "q", quality, "指定压缩质量, 压缩质量范围为 1(最低质量) - 100(最高质量), 质量越低, 图片大小越小.")
}

func run(cmd *cobra.Command, args []string) {
	if len(dirPath) == 0 {
		prompt := &survey.Input{
			Message: "📂 指定文件夹地址是什么?",
			Help:    "指定存放图片的文件夹, 支持在多级文件夹下获取图片.",
		}
		err := survey.AskOne(prompt, &dirPath)
		if err != nil {
			return
		}
	}

	if len(suffix) == 0 {
		prompt := &survey.Select{
			Message: "📌 需要压缩什么格式的图片?",
			Options: []string{"png", "jpeg", "jpg"},
			Default: "jpeg",
		}
		err := survey.AskOne(prompt, &suffix)
		if err != nil {
			return
		}
	}

	if len(quality) == 0 {
		prompt := &survey.Input{
			Message: "🔍 压缩质量预期是多大?",
			Help:    "压缩质量范围为 1(最低质量) - 100(最高质量), 质量越低, 图片大小越小.",
			Default: "20",
		}
		err := survey.AskOne(prompt, &quality)
		if err != nil {
			return
		}
	}

	if err := New(suffix, dirPath, quality); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
		return
	}
}