package main

import (
	"github.com/YuriyLisovskiy/sloc/src/parser"
	"github.com/YuriyLisovskiy/sloc/src/utils"
)

func main() {
//	println("Hello, sloc!")

	languages := []parser.Lang{
		{Name: "C++", FilesCount: 1, LinesCount: 135, BlankLinesCount: 10, CommentLinesCount: 12, CodeLinesCount: 113},
		{Name: "Html", FilesCount: 3, LinesCount: 390, BlankLinesCount: 103, CommentLinesCount: 56, CodeLinesCount: 231},
	}

	total := parser.Lang {Name: "Total"}
	for _, item := range languages {
		total.CodeLinesCount += item.CodeLinesCount
		total.CommentLinesCount += item.CommentLinesCount
		total.BlankLinesCount += item.BlankLinesCount
		total.LinesCount += item.LinesCount
		total.FilesCount += item.FilesCount
	}

	utils.OutputToStd(languages, total)

/*
	d, f, ff, j, err := args.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(d, f, ff, j)
*/
}
