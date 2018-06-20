package parser

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/utils"
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


)

func TestNormalizeLang(test *testing.T) {
	for _, td := range normalizeLangTestData {
		actual := NormalizeLang(td.input)
		if actual != td.expected {
			test.Errorf("parser.TestNormalizeLang: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestGetExt(test *testing.T) {
	for _, td := range getExtTestData {
		actual := GetExt(td.input)
		if actual != td.expected {
			test.Errorf("parser.TestGetExt: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestExtIsRecognized(test *testing.T) {
	for _, td := range extIsRecognizedTestData {
		actual := ExtIsRecognized(td.ext)
		if actual != td.expected {
			test.Errorf("parser.TestExtIsRecognized: expected %t, actual %t", td.expected, actual)
		}
	}
}

func TestGetFileNameFromPath(test *testing.T) {
	for _, td := range getFileNameFromPathTestData {
		actual := getFileNameFromPath(td.input)
		if actual != td.expected {
			test.Errorf("parser.TestGetFileNameFromPath: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestConcatLangs(test *testing.T) {
	for _, td := range concatLangsTestData {
		actual := utils.MergeLangs(td.current, td.lang)
		if actual != td.expected {
			test.Errorf("parser.TestConcatLangs: expected %s, actual %s", td.expected.ToString(), actual.ToString())
		}
	}
}

func TestSplitFile(test *testing.T) {
	for _, td := range splitFileTestData {
		actual := SplitFile(td.input)
		for i, obj := range actual {
			if obj != td.expected[i] {
				test.Errorf("parser.TestSplitFile: expected %s, actual %s", td.expected[i], obj)
			}
		}
	}
}
