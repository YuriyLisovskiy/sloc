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

func Parse() (string, string, []string, bool, error) {
	flag.Parse()
	multipleFiles := SplitMultFiles(*filesPtr)
	err := validate(*dirPtr, *filePtr, multipleFiles)
	if err != nil {
		println(err.Error())
	}
	return *dirPtr, *filePtr, multipleFiles, *jsonOutPtr, err
}
