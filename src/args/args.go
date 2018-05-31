package args

import (
	"flag"
	"strings"
)

func SplitMultFiles(filesStr string) []string {
	var result []string
	if len(filesStr) > 0 {
		files := strings.Split(strings.Join(strings.Fields(filesStr), " "), " ")
		for _, file := range files {
			result = append(result, strings.TrimSpace(file))
		}
	}
	return result
}

func Parse() (string, string, []string, bool, bool, bool, bool, error) {
	flag.Parse()
	parseOutPath()
	multipleFiles := SplitMultFiles(*filesPtr)
	var err error = nil
	isHelp := false
	if *h || *help {
		isHelp = true
		flag.Usage()
	} else {
		err = Validate(*dirPtr, *filePtr, multipleFiles)
	}
	return *dirPtr, *filePtr, multipleFiles, *jsonOutPtr, *xmlOutPtr, *ymlOutPtr, isHelp, err
}

func eraseExt(name string) string {
	if index := strings.LastIndex(name, "."); index != -1 {
		name = name[:index]
	}
	return name
}

func parseOutPath() {
	path := *OutPathPtr
	if path == "" || path == "." {
		path = "./"
	}
	if path != "./" {
		pathLen := len(path)
		if pathLen > 0 {
			if path[pathLen-1] != '/' {
				index := strings.LastIndex(path, "/")
				if index == -1 {
					DefaultOutFileName = eraseExt(path)
					path = "./"
				} else {
					DefaultOutFileName = eraseExt(path[index+1:])
					path = path[:index+1]
				}
			}
		}
	}
	*OutPathPtr = path
}
