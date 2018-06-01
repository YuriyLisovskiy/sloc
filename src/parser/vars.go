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
		"c",
		"cpp",
		"cs",
		"ccpph",
		"fs",
		"go",
		"groovy",
		"hs",
		"html",
		"java",
		"js",
		"json",
		"kt",
		"Makefile",
		"md",
		"py",
		"rb",
		"txt",
		"xml",
		"yml",
	}
	languageData = map[string]models.AvailableLang{
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
		"cs": {
			"C#",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"ccpph": {
			"C/C++ Header",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"fs": {
			"F#",
			models.Comment{"//", "\n"},
			models.Comment{"(*", "*)"},
		},
		"go": {
			"Go",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"groovy": {
			"Groovy",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"hs": {
			"Haskell",
			models.Comment{"--", "\n"},
			models.Comment{"{-", "-}"},
		},
		"html": {
			"Html",
			models.Comment{"<!--", "-->"},
			models.Comment{"<!--", "-->"},
		},
		"java": {
			"Java",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"js": {
			"JavaScript",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"json": {
			"Json",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"kt": {
			"Kotlin",
			models.Comment{"//", "\n"},
			models.Comment{"/*", "*/"},
		},
		"Makefile": {
			"Makefile",
			models.Comment{"#", "\n"},
			models.Comment{emptyString, emptyString},
		},
		"md": {
			"Markdown",
			models.Comment{"[//]: # ", "\n"},
			models.Comment{"<!--", "-->"},
		},
		"txt": {
			"Plain text",
			models.Comment{emptyString, emptyString},
			models.Comment{emptyString, emptyString},
		},
		"py": {
			"Python",
			models.Comment{"#", "\n"},
			models.Comment{"\"\"\"", "\"\"\""},
		},
		"rb": {
			"Ruby",
			models.Comment{"#", "\n"},
			models.Comment{"=begin", "=end"},
		},
		"xml": {
			"Xml",
			models.Comment{"<!--", "-->"},
			models.Comment{"<!--", "-->"},
		},
		"yml": {
			"Yaml",
			models.Comment{"#", "\n"},
			models.Comment{"#", "\n"},
		},
	}
)
