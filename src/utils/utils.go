package utils

import (
	"os"
	"fmt"
	"math"
	"errors"
	"strings"
	"strconv"
	"io/ioutil"
	"github.com/YuriyLisovskiy/sloc/src/models"
)

func GetTemplates(langs []models.Lang) (string, string, string) {
	indent := findIndentLen(langs)
	header := fmt.Sprintf(getFormatTemplate(itos(indent), "s")+"\n", "Language", "Files", "Lines", "Blank", "Comments", "Code", )
	line := strings.Join(make([]string, len(header)+1), "-") + "\n"
	return line, line + header + line, getFormatTemplate(itos(indent), "d") + "\n"

}

func getFormatTemplate(indent, dataType string) string {
	return " %-" + indent + "s%" + indent +
		dataType + "%" + indent +
		dataType + "%" + indent +
		dataType + "%" + indent +
		dataType + "%" + indent + dataType
}

func itos(val int) string {
	return strconv.Itoa(val)
}

func itol(val int) int {
	return len(itos(val))
}

func findMax(first, second int) int {
	return int(math.Max(float64(first), float64(second)))
}

func findFieldWithMaxLen(lang models.Lang) int {
	return findMax(
		len(lang.Name)+1, findMax(
			itol(lang.FilesCount), findMax(
				itol(lang.LinesCount), findMax(
					itol(lang.BlankLinesCount), findMax(
						itol(lang.CommentLinesCount), findMax(
							itol(lang.CodeLinesCount), 12),
					),
				),
			),
		),
	)
}

func findIndentLen(langs []models.Lang) int {
	max := 0
	for _, lang := range langs {
		newMax := findFieldWithMaxLen(lang)
		if max < newMax {
			max = newMax
		}
	}
	return max
}

func AppendLangData(lang models.Lang, formatString string) string {
	return fmt.Sprintf(
		formatString,
		lang.Name,
		lang.FilesCount,
		lang.LinesCount,
		lang.BlankLinesCount,
		lang.CommentLinesCount,
		lang.CodeLinesCount,
	)
}

func WriteToFile(path, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New(fmt.Sprintf("cannot create file '%s': %s", path, err.Error()))
	}
	defer file.Close()
	fmt.Fprintf(file, data)
	return nil
}

func CreateDir(path string) error {
	return os.MkdirAll(path,0777)
}

func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().IsDir()
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return true
	}
	return info.Mode().IsRegular()
}

func MergeLangs(current, lang models.Lang) models.Lang {
	return models.Lang{
		Name:              current.Name,
		CodeLinesCount:    current.CodeLinesCount + lang.CodeLinesCount,
		CommentLinesCount: current.CommentLinesCount + lang.CommentLinesCount,
		BlankLinesCount:   current.BlankLinesCount + lang.BlankLinesCount,
		LinesCount:        current.LinesCount + lang.LinesCount,
		FilesCount:        current.FilesCount + lang.FilesCount,
	}
}

func ReadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func MapToArray(langMap map[string]*models.Lang) []models.Lang {
	var result []models.Lang
	for _, value := range langMap {
		result = append(result, *value)
	}
	return result
}

func NormalizePath(path string, isDirectory bool) string {
	if !strings.HasPrefix(path, "./") && !strings.HasPrefix(path, "/") {
		path = "./" + path
	}
	if isDirectory && !strings.HasSuffix(path, "/") {
		path += "/"
	}
	return path
}
