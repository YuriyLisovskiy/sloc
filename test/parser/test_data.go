package parser

import (
	"github.com/YuriyLisovskiy/sloc/src/parser"
	"github.com/YuriyLisovskiy/sloc/src/models"
)

var (
	normalizeLangTestData = []struct {
		input    string
		expected string
	}{
		{"cs", "cs"},
		{"h", "ccpph"},
		{"cc", "cpp"},
		{"yaml", "yml"},
		{"htm", "html"},
	}

	getExtTestData = []struct {
		input    string
		expected string
	}{
		{"file.cs", "cs"},
		{"file.hpp", "hpp"},
		{"file.go", "go"},
	}

	extIsRecognizedTestData = []struct {
		ext       string
		expected  bool
	}{
		{"txt", true},
		{"c", true},
		{"sass", true},
		{"hpp", false},
	}

	getFileNameFromPathTestData = []struct {
		input    string
		expected string
	}{
		{"src/file.go", "file.go"},
		{"src/file/some/", "src/file/some/"},
		{"Makefile", "Makefile"},
		{"path/to/make/file/Makefile", "Makefile"},
	}

	concatLangsTestData = []struct {
		current  models.Lang
		lang     models.Lang
		expected models.Lang
	}{
		{
			models.Lang{
				Name:              "Current",
				FilesCount:        10,
				CodeLinesCount:    100,
				CommentLinesCount: 34,
				BlankLinesCount:   1,
				LinesCount:        234,
			},
			models.Lang{
				Name:              "Some Lang",
				FilesCount:        111,
				CodeLinesCount:    2300,
				CommentLinesCount: 234,
				BlankLinesCount:   1233,
				LinesCount:        4567,
			},
			models.Lang{
				Name:              "Current",
				FilesCount:        121,
				CodeLinesCount:    2400,
				CommentLinesCount: 268,
				BlankLinesCount:   1234,
				LinesCount:        4801,
			},
		},
	}

	splitFileTestData = []struct {
		input    string
		expected []string
	}{
		{
			"#include<iostream>\nvoid main()\n{\n\tstd::cout << \"Hello, World!\";\n}\n",
			[]string{"#include<iostream>", "void main()", "{", "\tstd::cout << \"Hello, World!\";", "}", ""},
		},
		{
			"if __name__ == '__main__':\n\tprint('Hello, World!')",
			[]string{"if __name__ == '__main__':", "\tprint('Hello, World!')"},
		},
	}

	parseLineTestData = []struct {
		line     string
		sc       string
		mc       string
		expected parser.Enum
	}{
		{"", "#", "#", parser.IsBlank},
		{"// This is commented line", "//", "/*", parser.IsSingleComment},
		{"/* This is single comment line */", "//", "/*", parser.IsMultiComment},
		{"/* This is multi comment line", "//", "/*", parser.IsMultiComment},
		{"def main():", "#", "\"\"\"", parser.IsCode},
	}

	parseMultiLineCommentTestData = []struct {
		lines      []string
		endComment string
		expected   int
	}{
		{[]string{"/* Hello,", "World!", "Comment */", "func main() {"}, "*/", 3},
		{[]string{"/* Hello World Comment */", "func main() {"}, "*/", 1},
		{[]string{"/* Hello World Comment */ func main() {"}, "*/", 1},
	}
)
