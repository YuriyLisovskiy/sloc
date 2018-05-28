package parser

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

var getExtTestData = []struct {
	input    string
	expected string
}{
	{"file.cs", "cs"},
	{"file.hpp", "hpp"},
	{"file.go", "go"},
}

func TestGetExtension(t *testing.T) {
	for _, td := range getExtTestData {
		actual := parser.GetExt(td.input)
		if actual != td.expected {
			t.Errorf("parser.GetExt(%s): expected %s, actual %s", td.input, td.expected, actual)
		}
	}
}
