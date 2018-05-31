package parser

import (
	"errors"
	"strings"
)

func ParseLine(line, singleComment, multiComment string) (Enum) {
	ln := strings.TrimSpace(line)
	singleCLength, multiCLength := len(singleComment), len(multiComment)
	if len(ln) > 0 {
		if singleComment != emptyString || multiComment != emptyString {
			if len(ln) >= singleCLength && ln[:singleCLength] == singleComment {
				return isSingleComment
			}
			if len(ln) >= multiCLength && ln[:multiCLength] == multiComment {
				return isMultiComment
			}
		}
	} else {
		return isBlank
	}
	return isCode
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
	if !ExtIsRecognized(ext, availableExtensions) {
		return "", Lang{}, errors.New("unrecognized file")
	}
	langData := languageData[ext]
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
		case isBlank:
			blankLines += 1
		case isSingleComment:
			commentLines += 1
		case isMultiComment:
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

func Parse(files []string) ([]Lang, Lang) {
	langMap, total := make(map[string]Lang), Lang{Name: "Total"}
	for _, file := range files {
		key, val, err := ParseFile(file)
		if err == nil {
			total = ConcatLangs(total, val)
			if _, ok := langMap[key]; ok {
				val = ConcatLangs(langMap[key], val)
			}
			langMap[key] = val
		}
	}
	var result []Lang
	for _, value := range langMap {
		result = append(result, value)
	}
	return result, total
}
