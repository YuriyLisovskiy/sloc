package parser

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

var normalizeLangTestData = []struct {
	input    string
	expected string
}{
	{"cs", "cs"},
	{"h", "c"},
	{"hpp", "cpp"},
	{"yaml", "yml"},
	{"htm", "html"},
}

func TestNormalizeLang(t *testing.T) {
	for _, td := range normalizeLangTestData {
		actual := parser.NormalizeLang(td.input)
		if actual != td.expected {
			t.Errorf("parser.NormalizeLang(%s): expected %s, actual %s", td.input, td.expected, actual)
		}
	}
}
