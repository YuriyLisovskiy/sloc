package parser

import (
	"os"
	"log"
	"strings"
	"io/ioutil"
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

func ReadDir(path string) []string {
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
		case isDir:
			dirs = append(dirs, path+pathName+"/")
		case isRegular:
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
		return isDir
	}
	return isRegular
}

func PrintStatus(current, max int) {
	barWidth := 50
	progress := (current * barWidth) / max
	print("Progress: [")
	for i := 0; i < barWidth; i++ {
		if i < progress {
			print("#")
		} else if i == progress {
			print(">")
		} else {
			print(" ")
		}
	}
	print("]", progress * 2, "%\r")
}
