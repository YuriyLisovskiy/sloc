package cli

import (
	"flag"
	"strings"
	"github.com/YuriyLisovskiy/sloc/src/utils"
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

func Parse(f, d, e []string) ([]string, []string, []string, error) {
	flag.Parse()
	if len(d) == 0 && len(f) == 0 {
		d = append(d, "./")
	}
	parseExclude(e)
	return f, d, e, Validate(f, d)
}

func parseExclude(excludeList []string) {
	if len(excludeList) > 0 {
		for i := 0; i < len(ExcludeList); i++ {
			isDir := false
			if utils.IsDirectory(ExcludeList[i]) {
				isDir = true
			}
			ExcludeList[i] = utils.NormalizePath(ExcludeList[i], isDir)
		}
	}
}
