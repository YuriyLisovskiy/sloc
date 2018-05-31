package utils

import "github.com/YuriyLisovskiy/sloc/src/parser"

var (
	getTemplatesTestData = []struct {
		langs          []parser.Lang
		expectedLine   string
		expectedHeader string
		expectedFormat string
	}{
		{
			[]parser.Lang{
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
		input    parser.Lang
		expected int
	}{
		{
			parser.Lang{
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
			parser.Lang{
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
			parser.Lang{
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
		input    []parser.Lang
		expected int
	}{
		{
			[]parser.Lang{
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
			[]parser.Lang{
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
		lang parser.Lang
		format string
		expected string
	}{
		{
			parser.Lang{
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
