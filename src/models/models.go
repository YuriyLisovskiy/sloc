package models

import "fmt"

type Comment struct {
	Start string
	End   string
}

type AvailableLang struct {
	Name              string
	SingleLineComment Comment
	MultiLineComment  Comment
}

type Lang struct {
	Name              string
	FilesCount        int
	LinesCount        int
	BlankLinesCount   int
	CommentLinesCount int
	CodeLinesCount    int
}

func (lang Lang)ToString() string {
	return fmt.Sprintf(
		"%s %d %d %d %d %d",
		lang.Name,
		lang.FilesCount,
		lang.LinesCount,
		lang.BlankLinesCount,
		lang.CommentLinesCount,
		lang.CodeLinesCount,
	)
}
