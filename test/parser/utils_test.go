package parser

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/utils"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func TestNormalizeLang(test *testing.T) {
	for _, td := range normalizeLangTestData {
		actual := parser.NormalizeLang(td.input)
		if actual != td.expected {
			test.Errorf("parser.TestNormalizeLang: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestGetExt(test *testing.T) {
	for _, td := range getExtTestData {
		actual := parser.GetExt(td.input)
		if actual != td.expected {
			test.Errorf("parser.TestGetExt: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestExtIsRecognized(test *testing.T) {
	for _, td := range extIsRecognizedTestData {
		actual := parser.ExtIsRecognized(td.ext, td.available)
		if actual != td.expected {
			test.Errorf("parser.TestGetExt: expected %t, actual %t", td.expected, actual)
		}
	}
}

func TestGetFileNameFromPath(test *testing.T) {
	for _, td := range getFileNameFromPathTestData {
		actual := parser.GetFileNameFromPath(td.input)
		if actual != td.expected {
			test.Errorf("parser.TestGetFileNameFromPath: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestConcatLangs(test *testing.T) {
	for _, td := range concatLangsTestData {
		actual := utils.ConcatLangs(td.current, td.lang)
		if actual != td.expected {
			test.Errorf("parser.TestConcatLangs: expected %s, actual %s", td.expected.ToString(), actual.ToString())
		}
	}
}

func TestSplitFile(test *testing.T) {
	for _, td := range splitFileTestData {
		actual := parser.SplitFile(td.input)
		for i, obj := range actual {
			if obj != td.expected[i] {
				test.Errorf("parser.TestSplitFile: expected %s, actual %s", td.expected[i], obj)
			}
		}
	}
}
