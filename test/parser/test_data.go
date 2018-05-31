package parser

import "github.com/YuriyLisovskiy/sloc/src/parser"

var (
	availableExts = []string{"txt", "go", "cpp", "sh"}
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
		available []string
		expected  bool
	}{
		{"txt", availableExts, true},
		{"c", availableExts, false},
		{"sh", availableExts, true},
		{"hpp", availableExts, false},
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
		current  parser.Lang
		lang     parser.Lang
		expected parser.Lang
	}{
		{
			parser.Lang{
				Name:              "Current",
				FilesCount:        10,
				CodeLinesCount:    100,
				CommentLinesCount: 34,
				BlankLinesCount:   1,
				LinesCount:        234,
			},
			parser.Lang{
				Name:              "Some Lang",
				FilesCount:        111,
				CodeLinesCount:    2300,
				CommentLinesCount: 234,
				BlankLinesCount:   1233,
				LinesCount:        4567,
			},
			parser.Lang{
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
		input string
		expected []string
	} {
		{
			"#include<iostream>\nvoid main()\n{\n\tstd::cout << \"Hello, World!\";\n}\n",
			[]string {"#include<iostream>", "void main()", "{", "\tstd::cout << \"Hello, World!\";", "}", ""},
		},
		{
			"if __name__ == '__main__':\n\tprint('Hello, World!')",
			[]string{"if __name__ == '__main__':", "\tprint('Hello, World!')"},
		},
	}
)
