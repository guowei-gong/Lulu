package main

import (
	"github.com/bimg/cmd/lulu/internal/image"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:     "lulu",
	Short:   "Lulu: 一个文件压缩处理的工具箱.",
	Long:    `Lulu: 一个文件压缩处理的工具箱.`,
	Version: "1.0",
}

func init() {
	rootCmd.AddCommand(image.CmdImage)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
