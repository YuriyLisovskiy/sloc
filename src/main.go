package main

import (
	"github.com/YuriyLisovskiy/sloc/src/args"
	"github.com/YuriyLisovskiy/sloc/src/utils"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func main() {
	dir, file, multipleFiles, _, err := args.Parse()
	if err != nil {
		return
	}
	var res []parser.Lang
	var total parser.Lang
	if dir != "" {
		res, total = parser.Parse(parser.ReadDir(dir))
	} else if file != "" {
		res, total = parser.Parse([]string{file})
	} else if len(multipleFiles) > 0 {
		res, total = parser.Parse(multipleFiles)
	}
	utils.OutputToStd(res, total)
}
