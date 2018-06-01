package parser

import "github.com/YuriyLisovskiy/sloc/src/models"

type Enum int

const (
	IsSingleComment Enum = iota
	IsMultiComment  Enum = iota
	IsBlank         Enum = iota
	IsCode          Enum = iota
	emptyString            = ""
)

var (
	availableExtensions = []string{
		"ccpph", "c", "cpp", "go", "java", "rb", "py", "js", "cs", "xml", "html", "json", "yml", "txt",
	}
	languageData = map[string]models.AvailableLang{
		"ccpph": {
			"C/C++ Header",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"c": {
			"C",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"cpp": {
			"C++",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"go": {
			"Go",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"java": {
			"Java",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"rb": {
			"Ruby",
			models.Comment{"#", "\n"},
			models.Comment{"=begin", "=end"},
		},
		"py": {
			"Python",
			models.Comment{"#", "\n"},
			models.Comment{"\"\"\"", "\"\"\""},
		},
		"js": {
			"JavaScript",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"cs": {
			"C#",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"xml": {
			"Xml",
			models.Comment{"<!--", "-->"},
			models.Comment{"<!--", "-->"},
		},
		"html": {
			"Html",
			models.Comment{"<!--", "-->"},
			models.Comment{"<!--", "-->"},
		},
		"json": {
			"Json",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"yml": {
			"Yaml",
			models.Comment{"#", "\n"},
			models.Comment{"#", "\n"},
		},
		"txt": {
			"Plain text",
			models.Comment{emptyString, emptyString},
			models.Comment{emptyString, emptyString},
		},
	}
)
