package parser

import (
	"log"
	"fmt"
	"errors"
	"strings"
	"io/ioutil"
	"github.com/YuriyLisovskiy/sloc/src/utils"
	"github.com/YuriyLisovskiy/sloc/src/models"
)

func parseLine(line, singleComment, multiComment string) (Enum) {
	ln := strings.TrimSpace(line)
	if len(ln) > 0 {
		if singleComment != emptyString {
			if strings.HasPrefix(ln, singleComment) {
				return IsSingleComment
			}
		}
		if multiComment != emptyString {
			if strings.HasPrefix(ln, multiComment) {
				return IsMultiComment
			}
		}
	} else {
		return IsBlank
	}
	return IsCode
}

func parseMultiLineComment(lines []string, endComment string) (int) {
	index := 1
	for i, line := range lines {
		if strings.HasSuffix(line, endComment) {
			index += i
			break
		} else if strings.Contains(line, endComment) {
			index += i - 1
			break
		}
	}
	return index
}

func parseSingleFile(file, ext string) (models.Lang, error) {
	langData := languageData[ext]
	content, err := utils.ReadFile(file)
	if err != nil {
		return models.Lang{}, err
	}
	lines := SplitFile(content)
	blankLines := 0
	commentLines := 0
	for i := 0; i < len(lines); i++ {
		lineType := parseLine(lines[i], langData.SingleLineComment.Start, langData.MultiLineComment.Start)
		switch lineType {
		case IsBlank:
			blankLines += 1
		case IsSingleComment:
			commentLines += 1
		case IsMultiComment:
			newIndex := parseMultiLineComment(lines[i:], langData.MultiLineComment.End)
			commentLines += newIndex
			i += newIndex - 1
		}
	}
	result := models.Lang{
		Name:              langData.Name,
		CommentLinesCount: commentLines,
		BlankLinesCount:   blankLines,
		CodeLinesCount:    len(lines) - (blankLines + commentLines),
		LinesCount:        len(lines),
		FilesCount:        1,
	}
	return result, nil
}

func ParseMultiple(files []string) ([]models.Lang, models.Lang) {
	langMap, total := make(map[string]*models.Lang), models.Lang{Name: "Total"}
	var subTotal models.Lang
	for _, file := range files {
		if !IsExcluded(file) {
			if utils.IsDirectory(file) {
				_, subTotal = ParseDirectory(file, langMap)
				total = utils.MergeLangs(total, subTotal)
			} else if utils.IsFile(file) {
				ext := NormalizeLang(GetExt(file))
				if ExtIsRecognized(ext) {
					val, err := parseSingleFile(file, ext)
					if err == nil {
						total = utils.MergeLangs(total, val)
						if _, ok := langMap[ext]; ok {
							val = utils.MergeLangs(*langMap[ext], val)
						}
						langMap[ext] = &val
					}
				}
			}
		}
	}
	return utils.MapToArray(langMap), total
}

func ParseFile(file string) (models.Lang, error) {
	if !IsExcluded(file) {
		ext := NormalizeLang(GetExt(file))
		if ExtIsRecognized(ext) {
			val, err := parseSingleFile(file, ext)
			if err == nil {
				return val, nil
			}
		}
		return models.Lang{}, errors.New(fmt.Sprintf("can't parse file '%s'", file))
	}
	return models.Lang{}, errors.New(fmt.Sprintf("file is excluded: '%s'", file))
}

func ParseDirectory(path string, langMap map[string]*models.Lang) ([]models.Lang, models.Lang) {
	path = utils.NormalizePath(path, true)
	readResult, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	if langMap == nil {
		langMap = make(map[string]*models.Lang)
	}
	total := models.Lang{Name: "Total"}
	var subTotal models.Lang
	for _, pathData := range readResult {
		newPath := path + pathData.Name()
		if utils.IsDirectory(newPath) {
			newPath = utils.NormalizePath(newPath, true)
			if !IsExcluded(newPath) {
				_, subTotal = ParseDirectory(newPath, langMap)
				total = utils.MergeLangs(total, subTotal)
			}
		} else if utils.IsFile(newPath) {
			newPath = utils.NormalizePath(newPath, false)
			if !IsExcluded(newPath) {
				ext := NormalizeLang(GetExt(newPath))
				if ExtIsRecognized(ext) {
					val, err := parseSingleFile(newPath, ext)
					if err == nil {
						total = utils.MergeLangs(total, val)
						if _, ok := langMap[ext]; ok {
							val = utils.MergeLangs(*langMap[ext], val)
						}
						langMap[ext] = &val
					}
				}
			}
		}
	}
	return utils.MapToArray(langMap), total
}
