package parser

import (
	"fmt"
	"strings"
)

func ParseLine(line, singleComment, multiComment string) (Enum) {
	ln := strings.TrimSpace(line)
	singleCLength, multiCLength := len(singleComment), len(multiComment)
	if len(ln) > 0 {
		if singleComment != "" || multiComment != "" {
			if len(ln) >= singleCLength && ln[:singleCLength] == singleComment {
				return IsSingleComment
			}
			if len(ln) >= multiCLength && ln[:multiCLength] == multiComment {
				return IsMultiComment
			}
		}
	} else {
		return IsBlank
	}
	return Code
}

func ParseMultiLineComment(lines []string, endComment string) (int) {
	index := 1
	for i, line := range lines {
		if strings.Contains(line, endComment) {
			index += i
			break
		}
	}
	return index
}

func ParseFile(file string) (string, Lang, error) {
	ext := NormalizeLang(GetExt(file))
	langData := LangData[ext]
	content, err := ReadFile(file)
	if err != nil {
		return "", Lang{}, err
	}
	lines := SplitFile(content)
	blankLines := 0
	commentLines := 0
	for i := 0; i < len(lines); i++ {
		lineType := ParseLine(lines[i], langData.SingleLineComment.Start, langData.MultiLineComment.Start)
		switch lineType {
		case IsBlank:
			blankLines += 1
		case IsSingleComment:
			commentLines += 1
		case IsMultiComment:
			newIndex := ParseMultiLineComment(lines[i:], langData.MultiLineComment.End)
			commentLines += newIndex
			i += newIndex
		}
	}
	result := Lang{
		Name:              langData.Name,
		CommentLinesCount: commentLines,
		BlankLinesCount:   blankLines,
		CodeLinesCount:    len(lines) - (blankLines + commentLines),
		LinesCount:        len(lines),
		FilesCount:        1,
	}
	return ext, result, nil
}

func Parse(files []string) ([]Lang, Lang, Lang) {
	langMap, total := make(map[string]Lang), Lang{Name: "Total"}
	for _, file := range files {
		key, val, err := ParseFile(file)
		if err == nil {
			total = ConcatLangs(total, val)
			if _, ok := langMap[key]; ok {
				existentElem := langMap[key]
				existentElem.FilesCount += val.FilesCount
				existentElem.LinesCount += val.LinesCount
				existentElem.BlankLinesCount += val.BlankLinesCount
				existentElem.CommentLinesCount += val.CommentLinesCount
				existentElem.CodeLinesCount += val.CodeLinesCount
				val = existentElem
			}
			langMap[key] = val
		} else {
			fmt.Println("parser.Parse: something went wrong with file '" + file + "'...")
		}
	}
	var result []Lang
	other := Lang{}
	if _, ok := langMap["other"]; ok {
		other = langMap["other"]
	}
	for _, value := range langMap {
		result = append(result, value)
	}
	return result, other, total
}
