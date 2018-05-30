package parser

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func TestNormalizeLang(t *testing.T) {
	for _, td := range normalizeLangTestData {
		actual := parser.NormalizeLang(td.input)
		if actual != td.expected {
			t.Errorf("parser.NormalizeLang(%s): expected %s, actual %s", td.input, td.expected, actual)
		}
	}
}

func TestGetExtension(t *testing.T) {
	for _, td := range getExtTestData {
		actual := parser.GetExt(td.input)
		if actual != td.expected {
			t.Errorf("parser.GetExt(%s): expected %s, actual %s", td.input, td.expected, actual)
		}
	}
}
