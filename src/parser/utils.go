package parser

import (
	"strings"
	"io/ioutil"
	"log"
	"os"
)

func NormalizeLang(ext string) string {
	switch ext {
	case "h":
		ext = "c"
	case "hpp":
		ext = "cpp"
	case "htm":
		ext = "html"
	case "yaml":
		ext = "yml"
	}
	return ext
}

func GetExt(fileName string) string {
	arr := strings.Split(GetFileNameFromPath(fileName), ".")
	res := "other"
	if len(arr) > 0 {
		ext := arr[len(arr)-1]
		for _, e := range AvailableExtensions {
			if e == ext {
				res = ext
				break
			}
		}
	}
	return res
}

func GetFileNameFromPath(path string) string {
	index := strings.LastIndex(path, "/")
	if index == -1 {
		return path
	}
	return path[index+1:]
}

func ConcatLangs(current, lang Lang) Lang {
	return Lang{
		Name:              current.Name,
		CodeLinesCount:    current.CodeLinesCount + lang.CodeLinesCount,
		CommentLinesCount: current.CommentLinesCount + lang.CommentLinesCount,
		BlankLinesCount:   current.BlankLinesCount + lang.BlankLinesCount,
		LinesCount:        current.LinesCount + lang.LinesCount,
		FilesCount:        current.FilesCount + lang.FilesCount,
	}
}

func SplitFile(content string) []string {
	return strings.Split(content, "\n")
}

func ReadFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ReadDir(path string) ([]string) {
	readResult, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	if path[len(path)-1] != '/' {
		path += "/"
	}
	var files, dirs []string
	for _, pathData := range readResult {
		pathName := pathData.Name()
		switch pathInfo(path + pathName + "/") {
		case IsDir:
			dirs = append(dirs, path+pathName+"/")
		case IsRegular:
			files = append(files, path+pathName)
		}
	}
	for _, dir := range dirs {
		files = append(files, ReadDir(dir)...)
	}
	return files
}

func pathInfo(path string) Enum {
	if _, err := os.Stat(path); err == nil {
		return IsDir
	}
	return IsRegular
}
