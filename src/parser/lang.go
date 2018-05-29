package parser

type Lang struct {
	Name              string
	FilesCount        int
	LinesCount        int
	BlankLinesCount   int
	CommentLinesCount int
	CodeLinesCount    int
}

type Comment struct {
	Start string
	End   string
}

type AvailableLang struct {
	Name              string
	SingleLineComment Comment
	MultiLineComment  Comment
}

type LangList map[string]Lang

func (l LangList) exists(key string) bool {
	_, ok := l[key]
	return ok
}

/* Hello

*/
