package parser

import (
	"fmt"
	"regexp"
	"strings"
	"io/ioutil"
)

type Enum int

const (
	IsSingleComment Enum = iota
	IsMultiComment  Enum = iota
	IsBlank         Enum = iota
	Else            Enum = iota
)

func GetExt(fileName string) string {
	arr := strings.Split(fileName, ".")
	res := "other"
	if len(arr) > 0 {
		res = arr[len(arr)-1]
	}
	return res
}

func ReadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func SplitFile(content string) []string {
	return strings.Split(content, "\n")
}

func ParseLine(line, singleComment, multiComment string) (Enum) {
	ln := strings.TrimSpace(line)
	if len(ln) > 0 {
		if singleComment != "" || multiComment != "" {
			byteLn := []byte(ln)
			matchRes, err := regexp.Match(singleComment, byteLn)
			if err == nil && matchRes {
				return IsSingleComment
			}
			matchRes, err = regexp.Match(multiComment, byteLn)
			if err == nil && matchRes {
				return IsMultiComment
			}
		}
	} else {
		return IsBlank
	}
	return Else
}

func ParseMultiLineComment(lines []string, endComment string) (int) {
	index := 1
	for i, line := range lines {
		matchRes, err := regexp.Match(endComment, []byte(line))
		if err == nil {
			if matchRes {
				index += i
				break
			}
		}
	}
	return index
}

func ParseFile(file string) (string, Lang, error) {
	langData := LangData["other"]
	ext := GetExt(file)
	if _, ok := LangData[GetExt(file)]; ok {
		langData = LangData[ext]
	}
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

func ConcatLangs(current, lang Lang) Lang {
	return Lang{
		CodeLinesCount:    current.CodeLinesCount + lang.CodeLinesCount,
		CommentLinesCount: current.CommentLinesCount + lang.CommentLinesCount,
		BlankLinesCount:   current.BlankLinesCount + lang.BlankLinesCount,
		LinesCount:        current.LinesCount + lang.LinesCount,
		FilesCount:        current.FilesCount + lang.FilesCount,
	}
}

func Parse(files []string) ([]Lang, Lang) {
	var langMap = make(LangList)
	var total Lang
	for _, f := range files {
		key, val, err := ParseFile(f)
		if err == nil {
			total = ConcatLangs(total, val)
			if langMap.exists(key) {
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
			fmt.Println("parser.Parse: something went wrong with file '" + f + "'...")
		}
	}
	var result []Lang
	for _, value := range langMap {
		result = append(result, value)
	}
	total.Name = "Total"
	return result, total
}
