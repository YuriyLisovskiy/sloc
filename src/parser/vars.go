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
	emptyComment = models.Comment{emptyString, emptyString}
	cLikeSComment = models.Comment{"//", "\n"}
	cLikeMComment = models.Comment{"/*", "*/"}
	pyLikeSComment = models.Comment{"#", "\n"}
	xmlLikeComment = models.Comment{"<!--", "-->"}
	sqlLikeSComment = models.Comment{"--", "\n"}
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
		"makefile",
		"md",
		"m",
		"py",
		"rb",
		"txt",
		"xml",
		"yml",
	}
	languageData = map[string]models.AvailableLang{
		"c": {
			"C",
			cLikeSComment,
			cLikeMComment,
		},
		"cpp": {
			"C++",
			cLikeSComment,
			cLikeMComment,
		},
		"cs": {
			"C#",
			cLikeSComment,
			cLikeMComment,
		},
		"ccpph": {
			"C/C++ Header",
			cLikeSComment,
			cLikeMComment,
		},
		"fs": {
			"F#",
			cLikeSComment,
			models.Comment{"(*", "*)"},
		},
		"go": {
			"Go",
			cLikeSComment,
			cLikeMComment,
		},
		"groovy": {
			"Groovy",
			cLikeSComment,
			cLikeMComment,
		},
		"hs": {
			"Haskell",
			sqlLikeSComment,
			models.Comment{"{-", "-}"},
		},
		"html": {
			"Html",
			xmlLikeComment,
			xmlLikeComment,
		},
		"java": {
			"Java",
			cLikeSComment,
			cLikeMComment,
		},
		"js": {
			"JavaScript",
			cLikeSComment,
			cLikeMComment,
		},
		"json": {
			"Json",
			cLikeSComment,
			cLikeMComment,
		},
		"kt": {
			"Kotlin",
			cLikeSComment,
			cLikeMComment,
		},
		"makefile": {
			"Makefile",
			pyLikeSComment,
			emptyComment,
		},
		"md": {
			"Markdown",
			models.Comment{"[//]: # ", "\n"},
			xmlLikeComment,
		},
		"m": {
			"Objective-C",
			cLikeSComment,
			cLikeMComment,
		},
		"txt": {
			"Plain text",
			emptyComment,
			emptyComment,
		},
		"pas": {
			"Pascal",
			models.Comment{"{", "}"},
			models.Comment{"{*", "*}"},
		},
		"pl": {
			"Perl",
			pyLikeSComment,
			models.Comment{"=begin", "=cut"},
		},
		"php": {
			"PHP",
			cLikeSComment,
			cLikeMComment,
		},
		"py": {
			"Python",
			pyLikeSComment,
			models.Comment{"\"\"\"", "\"\"\""},
		},
		"r": {
			"R",
			pyLikeSComment,
			emptyComment,
		},
		"rb": {
			"Ruby",
			pyLikeSComment,
			models.Comment{"=begin", "=end"},
		},
		"sass": {
			"Sass",
			cLikeSComment,
			cLikeMComment,
		},
		"sql": {
			"SQL",
			sqlLikeSComment,
			cLikeMComment,
		},
		"swift": {
			"Swift",
			cLikeSComment,
			cLikeMComment,
		},
		"xml": {
			"Xml",
			xmlLikeComment,
			xmlLikeComment,
		},
		"yml": {
			"Yaml",
			pyLikeSComment,
			pyLikeSComment,
		},
	}
)
