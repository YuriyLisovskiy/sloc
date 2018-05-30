package parser

type Enum int

const (
	isSingleComment Enum = iota
	isMultiComment  Enum = iota
	isBlank         Enum = iota
	isCode            Enum = iota
	isDir           Enum = iota
	isRegular       Enum = iota
)

var (
	availableExtensions = []string{
		"c", "cpp", "go", "java", "rb", "py", "js", "cs", "xml", "html", "json", "yml",
	}
	languageData = map[string]AvailableLang{
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
			"Go",
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
		"other": {
			"Other",
			Comment{"", ""},
			Comment{"", ""},
		},
	}
)
