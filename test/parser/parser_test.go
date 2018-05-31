package parser

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func TestParseLine(test *testing.T) {
	for _, td := range parseLineTestData {
		actual := parser.ParseLine(td.line, td.sc, td.mc)
		if actual != td.expected {
			test.Errorf("parser.TestNormalizeLang: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestParseMultiLineComment(test *testing.T) {
	for _, td := range parseMultiLineCommentTestData {
		actual := parser.ParseMultiLineComment(td.lines, td.endComment)
		if actual != td.expected {
			test.Errorf("parser.TestNormalizeLang: expected %d, actual %d", td.expected, actual)
		}
	}
}
