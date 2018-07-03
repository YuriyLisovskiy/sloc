package parser

import (
	"strings"
	"github.com/YuriyLisovskiy/sloc/src/utils"
	"github.com/YuriyLisovskiy/sloc/src/models"
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
	case "tesc", "tese", "geom", "frag", "comp":
		ext = "vert"
	case "handlebars":
		ext = "hbs"
	case "lidr":
		ext = "idr"
	case "hlean":
		ext = "lean"
	case "el", "lsp", "scm", "ss", "rkt":
		ext = "lisp"
	case "mli":
		ext = "ml"
	}
	return ext
}

func GetExt(fileName string) string {
	arr := strings.Split(getFileNameFromPath(fileName), ".")
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

func getFileNameFromPath(path string) string {
	index := strings.LastIndex(path, "/")
	if index == -1 {
		return path
	}
	pathLen := len(path)
	if pathLen > 0 {
		if path[pathLen-1] == '/' {
			return path
		}
	}
	return path[index+1:]
}

func SplitFile(content string) []string {
	return strings.Split(content, "\n")
}

func IsExcluded(path string) bool {
	for i, excluded := range ExcludeList {
		if utils.NormalizePath(excluded, utils.IsDirectory(excluded)) == path {
			switch i {
			case len(ExcludeList) - 1:
				ExcludeList = append(ExcludeList[:i], []string{}...)
			default:
				ExcludeList = append(ExcludeList[:i], ExcludeList[i+1:]...)
			}
			return true
		}
	}
	return false
}

func language(lang string, slc, mlc models.Comment) models.AvailableLang {
	return models.AvailableLang{Name: lang, SingleLineComment: slc, MultiLineComment: mlc}
}

func cStyle(lang string) models.AvailableLang {
	return language(lang, clangSComment, clangMComment)
}

func shStyle(lang string) models.AvailableLang {
	return language(lang, shSComment, noneComments)
}

func noComments(lang string) models.AvailableLang {
	return language(lang, noneComments, noneComments)
}

func htmlStyle(lang string) models.AvailableLang {
	return language(lang, noneComments, xmlComment)
}

func mlStyle(lang string) models.AvailableLang {
	return language(lang, noneComments, mlMComment)
}
