package utils

import (
	"fmt"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func OutputToStd(langs []parser.Lang, total parser.Lang) {
	if len(langs) > 0 {
		line, header, dataFormat := GetTemplates(langs)
		stdOut := header
		for _, lang := range langs {
			stdOut += AppendLangData(lang, dataFormat)
		}
		stdOut += line + AppendLangData(total, dataFormat)
		fmt.Println(stdOut)
	} else {
		fmt.Println("There is no any files to count")
	}
}

func OutputToJson(langs []parser.Lang, total parser.Lang) {

}
