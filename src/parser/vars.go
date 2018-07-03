package parser

import "github.com/YuriyLisovskiy/sloc/src/models"

type Enum int

const (
	IsSingleComment Enum = iota
	IsMultiComment
	IsBlank
	IsCode
	emptyString     = ""
)

var ExcludeList []string

var (
	noComments    = models.Comment{emptyString, emptyString}
	clangSComment = models.Comment{"//", "\n"}
	clangMComment = models.Comment{"/*", "*/"}
	shSComment    = models.Comment{"#", "\n"}
	xmlComment    = models.Comment{"<!--", "-->"}
	sqlSComment   = models.Comment{"--", "\n"}
	agdaMComment  = models.Comment{"{-", "-}"}
	aspSComment   = models.Comment{"'", "\n"}
	mlMComment   = models.Comment{"(*", "*)"}
)

var (
	languageData = map[string]models.AvailableLang{
		"as":       cStyleLang("ActionScript"),
		"c":        {"C", clangSComment, clangMComment},
		"cpp":      {"C++", clangSComment, clangMComment},
		"cs":       {"C#", clangSComment, clangMComment},
		"ccpph":    {"C/C++ Header", clangSComment, clangMComment},
		"fs":       {"F#", clangSComment, mlMComment},
		"go":       {"Go", clangSComment, clangMComment},
		"groovy":   {"Groovy", clangSComment, clangMComment},
		"hs":       {"Haskell", sqlSComment, agdaMComment},
		"html":     {"Html", xmlComment, xmlComment},
		"java":     {"Java", clangSComment, clangMComment},
		"js":       {"JavaScript", clangSComment, clangMComment},
		"json":     {"JSON", clangSComment, clangMComment},
		"kt":       {"Kotlin", clangSComment, clangMComment},
		"makefile": {"Makefile", shSComment, noComments},
		"md":       {"Markdown", models.Comment{"[//]: # ", "\n"}, xmlComment},
		"m":        {"Objective-C", clangSComment, clangMComment},
		"mm":       {"Objective-C", clangSComment, clangMComment},
		"txt":      {"Plain text", noComments, noComments},
		"pas":      {"Pascal", models.Comment{"{", "}"}, models.Comment{"{*", "*}"}},
		"pl":       {"Perl", shSComment, models.Comment{"=begin", "=cut"}},
		"php":      {"PHP", clangSComment, clangMComment},
		"py":       {"Python", shSComment, models.Comment{"\"\"\"", "\"\"\""}},
		"r":        {"R", shSComment, noComments},
		"rb":       {"Ruby", shSComment, models.Comment{"=begin", "=end"}},
		"sass":     {"Sass", clangSComment, clangMComment},
		"sql":      {"SQL", sqlSComment, clangMComment},
		"swift":    {"Swift", clangSComment, clangMComment},
		"xml":      {"Xml", noComments, xmlComment},
		"yml":      {"Yaml", shSComment, noComments},
		"ads":      {"Ada", sqlSComment, noComments},
		"agda":     {"Agda", sqlSComment, agdaMComment},
		"asp":      {"ASP", aspSComment, xmlComment},
		"aspx":     {"ASP.NET", aspSComment, xmlComment},
		"asm":      {"Assembly", shSComment, clangMComment},
		"in":       {"Autoconf", shSComment, noComments},
		"awk":      {"Awk", shSComment, noComments},
		"bat":      {"Batch", models.Comment{"REM", "\n"}, noComments},
		"sh":       {"Bourne Shell", shSComment, noComments},
		"csh":      {"C Shell", shSComment, noComments},
		"coffee":   {"CoffeeScript", shSComment, models.Comment{"###", "###"}},
		"cfc":      {"ColdFusionScript", clangSComment, clangMComment},
		"v":        {"Coq", noComments, mlMComment},
		"css":      {"CSS", clangSComment, clangMComment},
		"cu":       {"CUDA", clangSComment, clangMComment},
		"cuh":      {"CUDA Header", clangSComment, clangMComment},
		"d":        {"D", clangSComment, clangMComment},
		"dart":     {"Dart", clangSComment, clangMComment},
		"dts":      {"DeviceTree", clangSComment, clangMComment},
		"erl":      {"Erlang", models.Comment{"%", "\n"}, noComments},
		"forth":    {"Forth", models.Comment{"\\", "\n"}, models.Comment{"(", ")"}},
		"f":        {"FORTRAN Legacy", models.Comment{"!", "\n"}, noComments},
		"f03":      {"FORTRAN Modern", models.Comment{"!", "\n"}, noComments},
		"vert":     {"GLSL", clangSComment, clangMComment},
		"hbs":      {"Handlebars", xmlComment, models.Comment{"{{!", "}}"}},
		"hex":      {"Hex", noComments, noComments},
		"idr":      {"Idris", sqlSComment, agdaMComment},
		"ini":      {"INI", models.Comment{";", "\n"}, noComments},
		"ihex":     {"Intel Hex", noComments, noComments},
		"jai":      {"Jai", clangSComment, clangMComment},
		"jsx":      {"Jsx", clangSComment, clangMComment},
		"jl":       {"Julia", shSComment, models.Comment{"#=", "=#"}},
		"lean":     {"Lean", sqlSComment, models.Comment{"/-", "-/"}},
		"less":     {"Less", clangSComment, clangMComment},
		"lds":      {"LinkerScript", clangSComment, clangMComment},
		"lisp":     {"Lisp", models.Comment{";", "\n"}, models.Comment{"|#", "#|"}},
		"lua":      {"Lua", sqlSComment, models.Comment{"--[[", "]]"}},
		"mustache": {"Mustache", noComments, models.Comment{"{{!", "}}"}},
		"nim":      {"Nim", shSComment, noComments},
		"nix":      {"Nix", shSComment, models.Comment{"/*", "*/"}},
		"ml":       {"OCaml", noComments, mlMComment},
	}
)
