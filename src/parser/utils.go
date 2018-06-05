package parser

import (
	"strings"
	"github.com/YuriyLisovskiy/sloc/src/args"
)

func NormalizeLang(ext string) string {
	ext = strings.ToLower(ext)
	switch ext {
	case "h", "hpp", "hh", "hxx":
		ext = "ccpph"
	case "cc", "cxx", "c++", "cp":
		ext = "cpp"
	case "htm", "xhtml":
		ext = "html"
	case "yaml":
		ext = "yml"
	case "gvy", "gy", "gsh":
		ext = "groovy"
	case "fsi", "fsx", "fsscript":
		ext = "fs"
	case "lhs":
		ext = "hs"
	case "kts":
		ext = "kt"
	case "markdown", "mdown", "mkdn", "mkd", "mdwn", "mdtxt", "mdtext", "text", "rmd":
		ext = "md"
	case "mm":
		ext = "m"
	case "pp", "inc":
		ext = "pas"
	case "pm", "t", "pod":
		ext = "pl"
	case "phtml", "php3", "php4", "php5", "php7", "phps", "php-s":
		ext = "php"
	case "rdata", "rds", "rda":
		ext = "r"
	case "scss":
		ext = "sass"
	case "adb":
		ext = "ads"
	case "lagda":
		ext = "agda"
	case "cshtml", "vbhtml", "asax", "ascx", "ashx", "asmx", "axd", "dbml", "edmx", "resx", "svc":
		ext = "aspx"
	case "s":
		ext = "asm"
	case "btm", "cmd":
		ext = "bat"
	case "postcss":
		ext = "css"
	case "dtsi":
		ext = "dts"
	case "hrl":
		ext = "erl"
	case "4th", "fr", "frt", "fth", "f83", "fb", "fpm", "e4", "rx", "ft":
		ext = "forth"
	case "for", "ftn", "f77", "pfo":
		ext = "f"
	case "f08", "f90", "f95":
		ext = "f03"
	}
	return ext
}

func GetExt(fileName string) string {
	arr := strings.Split(GetFileNameFromPath(fileName), ".")
	res := fileName
	if len(arr) > 0 {
		res = arr[len(arr)-1]
	}
	return res
}

func ExtIsRecognized(ext string) bool {
	for key := range languageData {
		if key == ext {
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

func IsExcluded(path string) bool {
	for i, excluded := range args.ExcludeList {
		if excluded == path {
			switch i {
			case len(args.ExcludeList) - 1:
				args.ExcludeList = append(args.ExcludeList[:i], nil...)
			default:
				args.ExcludeList = append(args.ExcludeList[:i], args.ExcludeList[i+1:]...)
			}
			return true
		}
	}
	return false
}
