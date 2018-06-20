package parser

import "testing"

var (
	parseLineTestData = []struct {
		line     string
		sc       string
		mc       string
		expected Enum
	}{
		{"", "#", "#", IsBlank},
		{"// This is commented line", "//", "/*", IsSingleComment},
		{"/* This is single comment line */", "//", "/*", IsMultiComment},
		{"/* This is multi comment line", "//", "/*", IsMultiComment},
		{"def main():", "#", "\"\"\"", IsCode},
	}

	parseMultiLineCommentTestData = []struct {
		lines      []string
		endComment string
		expected   int
	}{
		{[]string{"/* Hello,", "World!", "Comment */", "func main() {"}, "*/", 3},
		{[]string{"/* Hello World Comment */", "func main() {"}, "*/", 1},
		{[]string{"/* Hello World Comment */ func main() {"}, "*/", 0},
	}
)

func TestParseLine(test *testing.T) {
	for _, td := range parseLineTestData {
		actual := parseLine(td.line, td.sc, td.mc)
		if actual != td.expected {
			test.Errorf("parser.TestNormalizeLang: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestParseMultiLineComment(test *testing.T) {
	for _, td := range parseMultiLineCommentTestData {
		actual := parseMultiLineComment(td.lines, td.endComment)
		if actual != td.expected {
			test.Errorf("parser.TestNormalizeLang: expected %d, actual %d", td.expected, actual)
		}
	}
}
