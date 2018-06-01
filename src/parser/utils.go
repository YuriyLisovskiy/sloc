package parser

import (
	"strings"
	"io/ioutil"
	"github.com/YuriyLisovskiy/sloc/src/models"
)

func NormalizeLang(ext string) string {
	switch ext {
	case "h", "hpp", "hh", "hxx":
		ext = "ccpph"
	case "cc", "cxx":
		ext = "cpp"
	case "htm", "xhtml":
		ext = "html"
	case "yaml":
		ext = "yml"
	}
	return ext
}

func GetExt(fileName string) string {
	arr := strings.Split(GetFileNameFromPath(fileName), ".")
	res := ""
	if len(arr) > 0 {
		return arr[len(arr)-1]
	}
	return res
}

func ExtIsRecognized(ext string, available []string) bool {
	for _, e := range available {
		if e == ext {
			return true
		}
	}
	return false
}

func GetFileNameFromPath(path string) string {
	index := strings.LastIndex(path, "/")
	if index == -1 {
		return path
	}
	pathLen := len(path)
	if pathLen > 0 {
		if path[pathLen - 1] == '/' {
			return path
		}
	}
	return path[index+1:]
}

func SplitFile(content string) []string {
	return strings.Split(content, "\n")
}

func readFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func mapToArray(langMap map[string]*models.Lang) []models.Lang {
	var result []models.Lang
	for _, value := range langMap {
		result = append(result, *value)
	}
	return result
}
