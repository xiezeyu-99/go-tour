package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiezeyu-99/go-programming-tour-book/tour/internal/word"
	"log"
	"strings"
)

const (
	ModeToUpper = iota + 1
	ModeToLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderScore
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部单词转大写",
	"2：全部单词转小写",
	"3：下划线单词转大写驼峰单词",
	"4：下划线单词转小写驼峰单词",
	"5：驼峰单词转为下划线单词",
}, "\n")

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeToUpper:
			content = word.ToUpper(str)
		case ModeToLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderScore:
			content = word.CamelCaseTounderScore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word 查看帮助文档")
		}
		log.Printf("输出结果：%s", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
