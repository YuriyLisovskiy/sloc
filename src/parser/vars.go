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
	noneComments    = models.Comment{emptyString, emptyString}
	clangSComment = models.Comment{"//", "\n"}
	clangMComment = models.Comment{"/*", "*/"}
	shSComment    = models.Comment{"#", "\n"}
	xmlComment    = models.Comment{"<!--", "-->"}
	sqlSComment   = models.Comment{"--", "\n"}
	agdaMComment  = models.Comment{"{-", "-}"}
	aspSComment   = models.Comment{"'", "\n"}
	mlMComment    = models.Comment{"(*", "*)"}
)

var (
	languageData = map[string]models.AvailableLang{
		"as":       cStyle("ActionScript"),
		"c":        cStyle("C"),
		"cpp":      cStyle("C++"),
		"cs":       cStyle("C#"),
		"ccpph":    cStyle("C/C++ Header"),
		"go":       cStyle("Go"),
		"groovy":   cStyle("Groovy"),
		"java":     cStyle("Java"),
		"js":       cStyle("JavaScript"),
		"json":     cStyle("JSON"),
		"kt":       cStyle("Kotlin"),
		"m":        cStyle("Objective-C"),
		"mm":       cStyle("Objective-C"),
		"php":      cStyle("PHP"),
		"sass":     cStyle("Sass"),
		"swift":    cStyle("Swift"),
		"cfc":      cStyle("ColdFusionScript"),
		"css":      cStyle("CSS"),
		"cu":       cStyle("CUDA"),
		"cuh":      cStyle("CUDA Header"),
		"d":        cStyle("D"),
		"dart":     cStyle("Dart"),
		"dts":      cStyle("DeviceTree"),
		"vert":     cStyle("GLSL"),
		"jai":      cStyle("Jai"),
		"jsx":      cStyle("Jsx"),
		"less":     cStyle("Less"),
		"lds":      cStyle("LinkerScript"),

		"txt":      noComments("Plain text"),
		"hex":      noComments("Hex"),
		"ihex":     noComments("Intel Hex"),

		"html":     htmlStyle("Html"),
		"xml":      htmlStyle("Xml"),

		"makefile": shStyle("Makefile"),
		"r":        shStyle("R"),
		"yml":      shStyle("Yaml"),
		"in":       shStyle("Autoconf"),
		"awk":      shStyle("Awk"),
		"sh":       shStyle("Bourne Shell"),
		"csh":      shStyle("C Shell"),
		"nim":      shStyle("Nim"),

		"v":        mlStyle("Coq"),
		"ml":       mlStyle("OCaml"),

		"md":       language("Markdown", models.Comment{"[//]: # ", "\n"}, xmlComment),
		"pas":      language("Pascal", models.Comment{"{", "}"}, models.Comment{"{*", "*}"}),
		"pl":       language("Perl", shSComment, models.Comment{"=begin", "=cut"}),
		"py":       language("Python", shSComment, models.Comment{"\"\"\"", "\"\"\""}),
		"rb":       language("Ruby", shSComment, models.Comment{"=begin", "=end"}),
		"sql":      language("SQL", sqlSComment, clangMComment),
		"ads":      language("Ada", sqlSComment, noneComments),
		"agda":     language("Agda", sqlSComment, agdaMComment),
		"asp":      language("ASP", aspSComment, xmlComment),
		"aspx":     language("ASP.NET", aspSComment, xmlComment),
		"asm":      language("Assembly", shSComment, clangMComment),
		"bat":      language("Batch", models.Comment{"REM", "\n"}, noneComments),
		"coffee":   language("CoffeeScript", shSComment, models.Comment{"###", "###"}),
		"fs":       language("F#", clangSComment, mlMComment),
		"hs":       language("Haskell", sqlSComment, agdaMComment),
		"erl":      language("Erlang", models.Comment{"%", "\n"}, noneComments),
		"forth":    language("Forth", models.Comment{"\\", "\n"}, models.Comment{"(", ")"}),
		"f":        language("FORTRAN Legacy", models.Comment{"!", "\n"}, noneComments),
		"f03":      language("FORTRAN Modern", models.Comment{"!", "\n"}, noneComments),
		"hbs":      language("Handlebars", xmlComment, models.Comment{"{{!", "}}"}),
		"idr":      language("Idris", sqlSComment, agdaMComment),
		"ini":      language("INI", models.Comment{";", "\n"}, noneComments),
		"jl":       language("Julia", shSComment, models.Comment{"#=", "=#"}),
		"lean":     language("Lean", sqlSComment, models.Comment{"/-", "-/"}),
		"lisp":     language("Lisp", models.Comment{";", "\n"}, models.Comment{"|#", "#|"}),
		"lua":      language("Lua", sqlSComment, models.Comment{"--[[", "]]"}),
		"mustache": language("Mustache", noneComments, models.Comment{"{{!", "}}"}),
		"nix":      language("Nix", shSComment, models.Comment{"/*", "*/"}),
	}
)
