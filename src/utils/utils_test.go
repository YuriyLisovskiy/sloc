package utils

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/models"
)

var (
	getTemplatesTestData = []struct {
		langs          []models.Lang
		expectedLine   string
		expectedHeader string
		expectedFormat string
	}{
		{
			[]models.Lang{
				{
					Name:              "Name",
					CodeLinesCount:    10875,
					CommentLinesCount: 13434,
					BlankLinesCount:   34434,
					LinesCount:        111,
					FilesCount:        1343434,
				},
			},
			"--------------------------------------------------------------------------\n",
			"--------------------------------------------------------------------------\n Language           Files       Lines       Blank    Comments        Code\n--------------------------------------------------------------------------\n",
			" %-12s%12d%12d%12d%12d%12d\n",
		},
	}

	getFormatTemplateTestData = []struct {
		indent   string
		dataType string
		expected string
	}{
		{"5", "d", " %-5s%5d%5d%5d%5d%5d"},
		{"100", "s", " %-100s%100s%100s%100s%100s%100s"},
		{"33", "t", " %-33s%33t%33t%33t%33t%33t"},
	}

	itosTestData = []struct {
		input    int
		expected string
	}{
		{12345, "12345"},
		{0, "0"},
		{1111111, "1111111"},
		{234, "234"},
	}

	itolTestData = []struct {
		input    int
		expected int
	}{
		{12345, 5},
		{0, 1},
		{1111111, 7},
		{234, 3},
	}

	findMaxTestData = []struct {
		first    int
		second   int
		expected int
	}{
		{100, 45, 100},
		{11, 3, 11},
		{0, 1, 1},
		{0, 0, 0},
	}

	findFieldWithMaxLenTestData = []struct {
		input    models.Lang
		expected int
	}{
		{
			models.Lang{
				Name:              "Name",
				CodeLinesCount:    10875,
				CommentLinesCount: 13434,
				BlankLinesCount:   34434,
				LinesCount:        111,
				FilesCount:        1343434,
			},
			12,
		},
		{
			models.Lang{
				Name:              "Some interesting language",
				CodeLinesCount:    10875,
				CommentLinesCount: 13434,
				BlankLinesCount:   34434,
				LinesCount:        111,
				FilesCount:        1343434,
			},
			26,
		},
		{
			models.Lang{
				Name:              "Some language",
				CodeLinesCount:    31083456789075,
				CommentLinesCount: 13434,
				BlankLinesCount:   34434,
				LinesCount:        111,
				FilesCount:        1343434,
			},
			14,
		},
	}

	findIndentLenTestData = []struct {
		input    []models.Lang
		expected int
	}{
		{
			[]models.Lang{
				{
					Name:              "Name",
					CodeLinesCount:    10875,
					CommentLinesCount: 13434,
					BlankLinesCount:   34434,
					LinesCount:        111,
					FilesCount:        1343434,
				},
				{
					Name:              "Some interesting language",
					CodeLinesCount:    10875,
					CommentLinesCount: 13434,
					BlankLinesCount:   34434,
					LinesCount:        111,
					FilesCount:        1343434,
				},
				{
					Name:              "Some language",
					CodeLinesCount:    31083456789075,
					CommentLinesCount: 13434,
					BlankLinesCount:   34434,
					LinesCount:        111,
					FilesCount:        1343434,
				},
			},
			26,
		},
		{
			[]models.Lang{
				{
					Name:              "Name",
					CodeLinesCount:    10875,
					CommentLinesCount: 13434,
					BlankLinesCount:   34434,
					LinesCount:        111,
					FilesCount:        1343434,
				},
				{
					Name:              "Language",
					CodeLinesCount:    10875,
					CommentLinesCount: 13434,
					BlankLinesCount:   34434,
					LinesCount:        111,
					FilesCount:        1343434,
				},
				{
					Name:              "Some language",
					CodeLinesCount:    31083456789075,
					CommentLinesCount: 13434,
					BlankLinesCount:   34434,
					LinesCount:        111,
					FilesCount:        1343434,
				},
			},
			14,
		},
	}

	appendLangDataTestData = []struct {
		lang models.Lang
		format string
		expected string
	}{
		{
			models.Lang{
				Name:              "Name",
				CodeLinesCount:    10875,
				CommentLinesCount: 13434,
				BlankLinesCount:   34434,
				LinesCount:        111,
				FilesCount:        1343434,
			},
			" %-12s%12d%12d%12d%12d%12d\n",
			" Name             1343434         111       34434       13434       10875\n",
		},
	}
)

func TestGetTemplates(test *testing.T) {
	for _, td := range getTemplatesTestData {
		actualLine, actualHeader, actualFormat := GetTemplates(td.langs)
		if actualLine != td.expectedLine {
			test.Errorf("utils.TestGetTemplates: expected %s, actual %s", td.expectedLine, actualLine)
		}
		if actualHeader != td.expectedHeader {
			test.Errorf("utils.TestGetTemplates: expected %s, actual %s", td.expectedHeader, actualHeader)
		}
		if actualFormat != td.expectedFormat {
			test.Errorf("utils.TestGetTemplates: expected %s, actual %s", td.expectedFormat, actualFormat)
		}
	}
}

func TestGetFormatTemplate(test *testing.T) {
	for _, td := range getFormatTemplateTestData {
		actual := getFormatTemplate(td.indent, td.dataType)
		if actual != td.expected {
			test.Errorf("utils.TestGetFormatTemplate: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestItos(test *testing.T) {
	for _, td := range itosTestData {
		actual := itos(td.input)
		if actual != td.expected {
			test.Errorf("utils.TestItos: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestItol(test *testing.T) {
	for _, td := range itolTestData {
		actual := itol(td.input)
		if actual != td.expected {
			test.Errorf("utils.TestItol: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestFindMax(test *testing.T) {
	for _, td := range findMaxTestData {
		actual := findMax(td.first, td.second)
		if actual != td.expected {
			test.Errorf("utils.TestFindMax: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestFindFieldWithMaxLen(test *testing.T) {
	for _, td := range findFieldWithMaxLenTestData {
		actual := findFieldWithMaxLen(td.input)
		if actual != td.expected {
			test.Errorf("utils.TestFindFieldWithMaxLen: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestFindIndentLen(test *testing.T) {
	for _, td := range findIndentLenTestData {
		actual := findIndentLen(td.input)
		if actual != td.expected {
			test.Errorf("utils.TestFindIndentLen: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestAppendLangData(test *testing.T) {
	for _, td := range appendLangDataTestData {
		actual := AppendLangData(td.lang, td.format)
		if actual != td.expected {
			test.Errorf("utils.TestAppendLangData: expected %s, actual %s", td.expected, actual)
		}
	}
}
