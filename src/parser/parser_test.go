package parser

import "testing"

var (
	parseLineTestData = []struct {
		line     string
		sc       string
		mcs       string
		mce       string
		expected Enum
	}{
		{"", "#", "#", "\n", IsBlank},
		{"// This is commented line", "//", "/*", "*/", IsSingleComment},
		{"/* This is single comment line */", "//", "/*", "*/", IsSingleComment},
		{"/* This is multi comment line", "//", "/*", "*/", IsMultiComment},
		{"def main():", "#", "\"\"\"", "\"\"\"", IsCode},
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

func TestParseLine(test *testing.T) {
	for _, td := range parseLineTestData {
		actual := parseLine(td.line, td.sc, td.mcs, td.mce)
		if actual != td.expected {
			test.Errorf("parser.TestParseLine: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestParseMultiLineComment(test *testing.T) {
	for _, td := range parseMultiLineCommentTestData {
		actual := parseMultiLineComment(td.lines, td.endComment)
		if actual != td.expected {
			test.Errorf("parser.TestParseMultiLineComment: expected %d, actual %d", td.expected, actual)
		}
	}
}
