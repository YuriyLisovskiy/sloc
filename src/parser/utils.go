package parser

var LangData = map[string]AvailableLang{
	"c": {
		"C",
		Comment{"//", "\n"},
		Comment{"/*", "*/"},
	},
	"cpp": {
		"C++",
		Comment{"//", "\n"},
		Comment{"/*", "*/"},
	},
	"go": {
		"Golang",
		Comment{"//", "\n"},
		Comment{"/*", "*/"},
	},
	"java": {
		"Java",
		Comment{"//", "\n"},
		Comment{"/*", "*/"},
	},
	"rb": {
		"Ruby",
		Comment{"#", "\n"},
		Comment{"=begin", "=end"},
	},
	"py": {
		"Python",
		Comment{"#", "\n"},
		Comment{"\"\"\"", "\"\"\""},
	},
	"js": {
		"JavaScript",
		Comment{"//", "\n"},
		Comment{"/*", "*/"},
	},
	"cs": {
		"C#",
		Comment{"//", "\n"},
		Comment{"/*", "*/"},
	},
	"xml": {
		"Xml",
		Comment{"<!--", "-->"},
		Comment{"<!--", "-->"},
	},
	"html": {
		"Html",
		Comment{"<!--", "-->"},
		Comment{"<!--", "-->"},
	},
	"json": {
		"Json",
		Comment{"//", "\n"},
		Comment{"/*", "*/"},
	},
	"yml": {
		"Yaml",
		Comment{"#", "\n"},
		Comment{"#", "\n"},
	},
}

func NormalizeLang(ext string) string {
	switch ext {
	case "h":
		ext = "c"
	case "hpp":
		ext = "cpp"
	case "htm":
		ext = "html"
	case "yaml":
		ext = "yml"
	}
	return ext
}
