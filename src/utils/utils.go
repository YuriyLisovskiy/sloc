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
	indent := FindIndentLen(langs)
	header := fmt.Sprintf(GetFormatTemplate(Itos(indent), "s")+"\n", "Language", "Files", "Lines", "Blank", "Comments", "Code", )
	line := strings.Join(make([]string, len(header)+1), "-") + "\n"
	return line, line + header + line, GetFormatTemplate(Itos(indent), "d") + "\n"

}

func GetFormatTemplate(indent, dataType string) string {
	return " %-" + indent + "s%" + indent +
		dataType + "%" + indent +
		dataType + "%" + indent +
		dataType + "%" + indent +
		dataType + "%" + indent + dataType
}

func Itos(val int) string {
	return strconv.Itoa(val)
}

func Itol(val int) int {
	return len(Itos(val))
}

func FindMax(first, second int) int {
	return int(math.Max(float64(first), float64(second)))
}

func FindFieldWithMaxLen(lang models.Lang) int {
	return FindMax(
		len(lang.Name)+1, FindMax(
			Itol(lang.FilesCount), FindMax(
				Itol(lang.LinesCount), FindMax(
					Itol(lang.BlankLinesCount), FindMax(
						Itol(lang.CommentLinesCount), FindMax(
							Itol(lang.CodeLinesCount), 12),
					),
				),
			),
		),
	)
}

func FindIndentLen(langs []models.Lang) int {
	max := 0
	for _, lang := range langs {
		newMax := FindFieldWithMaxLen(lang)
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
