package parser

type Enum int

const (
	IsSingleComment Enum = iota
	IsMultiComment  Enum = iota
	IsBlank         Enum = iota
	Code            Enum = iota
	IsDir           Enum = iota
	IsRegular       Enum = iota
	Err             Enum = iota
)

var (
	AvailableExtensions = []string{
		"c", "cpp", "go", "java", "rb", "py", "js", "cs", "xml", "html", "json", "yml",
	}
	LangData = map[string]AvailableLang{
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
