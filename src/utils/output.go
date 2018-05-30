package utils

import (
	"fmt"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func appendLangData(lang parser.Lang) string {
	return fmt.Sprintf(
		stdOutFormatData,
		lang.Name,
		lang.FilesCount,
		lang.LinesCount,
		lang.BlankLinesCount,
		lang.CommentLinesCount,
		lang.CodeLinesCount,
	)
}

func OutputToStd(langs []parser.Lang, total parser.Lang, other *parser.Lang) {
	if len(langs) > 0 {
		stdOut := stdOutHeader
		for _, lang := range langs {
			stdOut += appendLangData(lang)
		}
		if other != nil {
			stdOut += stdOutLine + appendLangData(*other)
		}
		stdOut += stdOutLine + appendLangData(total)
		fmt.Println(stdOut)
	} else {
		fmt.Println("There is no any files to count")
	}
}

func OutputToJson(langs []parser.Lang, total parser.Lang) {

}
