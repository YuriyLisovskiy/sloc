package parser

import "strings"

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
