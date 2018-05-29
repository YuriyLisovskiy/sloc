package parser

var LangData = map[string]AvailableLang{
	"c": {
		"C",
		Comment{"\\s*//", "\n"},
		Comment{"\\s*/\\*", "\\*/\\s*"},
	},
	"cpp": {
		"C++",
		Comment{"\\s*//", "\n"},
		Comment{"\\s*/\\*", "\\*/\\s*"},
	},
	"go": {
		"Go",
		Comment{"\\s*//", "\n"},
		Comment{"\\s*/\\*", "\\*/\\s*"},
	},
	"java": {
		"Java",
		Comment{"\\s*//", "\n"},
		Comment{"\\s*/\\*", "\\*/\\s*"},
	},
	"rb": {
		"Ruby",
		Comment{"\\s*#", "\n"},
		Comment{"\\s*=begin", "=end\\s*"},
	},
	"py": {
		"Python",
		Comment{"\\s*#", "\n"},
		Comment{"\\s*\"\"\"", "\"\"\"\\s*"},
	},
	"js": {
		"JavaScript",
		Comment{"\\s*//", "\n"},
		Comment{"\\s*/\\*", "\\*/\\s*"},
	},
	"cs": {
		"C#",
		Comment{"\\s*//", "\n"},
		Comment{"\\s*/\\*", "\\*/\\s*"},
	},
	"xml": {
		"Xml",
		Comment{"\\s*<!--", "-->\\s*"},
		Comment{"\\s*<!--", "-->\\s*"},
	},
	"html": {
		"Html",
		Comment{"\\s*<!--", "-->\\s*"},
		Comment{"\\s*<!--", "-->\\s*"},
	},
	"json": {
		"Json",
		Comment{"\\s*//", "\n"},
		Comment{"\\s*/\\*", "\\*/\\s*"},
	},
	"yml": {
		"Yaml",
		Comment{"\\s*#", "\n"},
		Comment{"\\s*#", "\n"},
	},
	"other": {
		"Other",
		Comment{"", ""},
		Comment{"", ""},
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
