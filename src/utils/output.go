package utils

import (
	"fmt"
	"strings"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

var (
	stdOutLine = strings.Join(make([]string, 61), "-") + "\n"
	stdOutHeader = stdOutLine + "Language\tFiles\tLines\tBlank\tComments\tCode\n" + stdOutLine
	stdOutFormatData = "%s\t\t%d\t%d\t%d\t%d\t\t%d\n"
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

func OutputToStd(langs []parser.Lang, total parser.Lang, other parser.Lang, containsOther bool) {
	if len(langs) > 0 {
		stdOut := stdOutHeader
		for _, lang := range langs {
			stdOut += appendLangData(lang)
		}
		if containsOther {
			stdOut += stdOutLine + appendLangData(other)
		}
		stdOut += stdOutLine + appendLangData(total)
		fmt.Println(stdOut)
	} else {
		fmt.Println("There is no any files to count")
	}
}

func OutputToJson(langs []parser.Lang, total parser.Lang) {

}
