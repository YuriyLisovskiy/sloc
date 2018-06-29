package cli

import (
	"strings"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func splitMultFiles(filesStr string) []string {
	var result []string
	if len(filesStr) > 0 {
		files := strings.Split(strings.Join(strings.Fields(filesStr), " "), " ")
		for _, file := range files {
			result = append(result, strings.TrimSpace(file))
		}
	}
	return result
}

func Parse(f, d, m, e string) (string, string, []string, error) {
	if len(d) == 0 && len(f) == 0 && len(m) == 0 {
		d = "./"
	}
	parser.ExcludeList = splitMultFiles(e)
	d, f, multiple := parseExcluded(d, f, splitMultFiles(m))
	return d, f, multiple, Validate(d, f, m)
}

func parseExcluded(dir, file string, multiple []string) (string, string, []string) {
	if parser.IsExcluded(dir) {
		dir = ""
	}
	if parser.IsExcluded(file) {
		file = ""
	}
	var multipleResult []string
	for _, path := range multiple {
		if !parser.IsExcluded(path) {
			multipleResult = append(multipleResult, path)
		}
	}
	return dir, file, multipleResult
}
