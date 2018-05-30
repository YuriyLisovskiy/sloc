package args

import (
	"flag"
	"strings"
)

func SplitMultFiles(filesStr string) []string {
	files := strings.Split(strings.Join(strings.Fields(filesStr), " "), " ")
	var result []string
	for _, file := range files {
		result = append(result, strings.TrimSpace(file))
	}
	return result
}

func Parse() (string, string, []string, bool, error) {
	flag.Parse()
	multipleFiles := SplitMultFiles(*filesPtr)
//	err := validate(*dirPtr, *filePtr, multipleFiles)
	return *dirPtr, *filePtr, multipleFiles, *jsonOutPtr, nil
}
