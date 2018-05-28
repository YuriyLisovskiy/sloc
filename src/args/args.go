package args

import (
	"flag"
	"strings"
)

var (
	dirFlag = "d"
	fileFlag = "f"
	filesFlag = "mf"
	dirPtr = flag.String(dirFlag, "", "count lines of all files in directory")
	filePtr = flag.String(fileFlag, "", "count lines of one file")
	filesPtr = flag.String(fileFlag, "", "count lines of multiple files")
	jsonOutPtr = flag.Bool("json", false, "generate json result (default is std output)")
)

func splitMultFiles(filesStr string) []string {
	files := strings.Split(filesStr, " ")
	var result []string
	for _, file := range files {
		result = append(result, strings.TrimSpace(file))
	}
	return result
}

func Get() (string, string, []string, bool, error) {
	flag.Parse()
	multipleFiles := splitMultFiles(*filesPtr)
	err := validate(*dirPtr, *filePtr, multipleFiles)
	return *dirPtr, *filePtr, multipleFiles, *jsonOutPtr, err
}
