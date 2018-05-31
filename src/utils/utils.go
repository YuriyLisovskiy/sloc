package utils

import (
	"fmt"
	"math"
	"strings"
	"strconv"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func GetTemplates(langs []parser.Lang) (string, string, string) {
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

func FindFieldWithMaxLen(lang parser.Lang) int {
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

func FindIndentLen(langs []parser.Lang) int {
	max := 0
	for _, lang := range langs {
		newMax := FindFieldWithMaxLen(lang)
		if max < newMax {
			max = newMax
		}
	}
	return max
}

func AppendLangData(lang parser.Lang, formatString string) string {
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
