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

func ParseLine(line, singleComment, multiComment string) (Enum) {
	ln := strings.TrimSpace(line)
	singleCLength, multiCLength := len(singleComment), len(multiComment)
	if len(ln) > 0 {
		if singleComment != emptyString || multiComment != emptyString {
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
	return IsCode
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

func ParseMultiple(files []string) ([]models.Lang, models.Lang) {
	langMap, total := make(map[string]*models.Lang), models.Lang{Name: "Total"}
	var subTotal models.Lang
	for _, file := range files {
		if !IsExcluded(file) {
			if utils.IsDirectory(file) {
				_, subTotal = ParseDirectory(file, langMap)
				total = utils.ConcatLangs(total, subTotal)
			} else if utils.IsFile(file) {
				ext := NormalizeLang(GetExt(file))
				if ExtIsRecognized(ext, availableExtensions) {
					val, err := parseSingleFile(file, ext)
					if err == nil {
						total = utils.ConcatLangs(total, val)
						if _, ok := langMap[ext]; ok {
							val = utils.ConcatLangs(*langMap[ext], val)
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
		if ExtIsRecognized(ext, availableExtensions) {
			val, err := parseSingleFile(file, ext)
			if err == nil {
				return val, nil
			}
		}
		return models.Lang{}, errors.New(fmt.Sprintf("can't parse file '%s'", file))
	}
	return models.Lang{}, errors.New(fmt.Sprintf("file is excluded: '%s'", file))
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

func ParseDirectory(path string, langMap map[string]*models.Lang) ([]models.Lang, models.Lang) {
	if path[len(path)-1] != '/' {
		path += "/"
	}
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
		pathName := pathData.Name()
		if utils.IsDirectory(path + pathName + "/") {
			if !IsExcluded(path + pathName + "/") {
				_, subTotal = ParseDirectory(path+pathName+"/", langMap)
				total = utils.ConcatLangs(total, subTotal)
			}
		} else if utils.IsFile(path + pathName) {
			if !IsExcluded(path + pathName) {
				ext := NormalizeLang(GetExt(pathName))
				if ExtIsRecognized(ext, availableExtensions) {
					val, err := parseSingleFile(path+pathName, ext)
					if err == nil {
						total = utils.ConcatLangs(total, val)
						if _, ok := langMap[ext]; ok {
							val = utils.ConcatLangs(*langMap[ext], val)
						}
						langMap[ext] = &val
					}
				}
			}
		}
	}
	return utils.MapToArray(langMap), total
}
