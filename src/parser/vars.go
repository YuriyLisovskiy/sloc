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
	noneComment   = models.Comment{emptyString, emptyString}
	clangSComment = models.Comment{"//", "\n"}
	clangMComment = models.Comment{"/*", "*/"}
	shSComment    = models.Comment{"#", "\n"}
	xmlComment    = models.Comment{"<!--", "-->"}
	sqlSComment   = models.Comment{"--", "\n"}
	agdaMComment  = models.Comment{"{-", "-}"}
	aspSComment   = models.Comment{"'", "\n"}
	mlMComment    = models.Comment{"(*", "*)"}
	lispMComment  = models.Comment{"|#", "#|"}
)

var (
	languageData = map[string]models.AvailableLang{
		"as":     cStyle("ActionScript"),
		"c":      cStyle("C"),
		"cpp":    cStyle("C++"),
		"cs":     cStyle("C#"),
		"ccpph":  cStyle("C/C++ Header"),
		"go":     cStyle("Go"),
		"groovy": cStyle("Groovy"),
		"java":   cStyle("Java"),
		"js":     cStyle("JavaScript"),
		"json":   cStyle("JSON"),
		"kt":     cStyle("Kotlin"),
		"m":      cStyle("Objective-C"),
		"mm":     cStyle("Objective-C"),
		"php":    cStyle("PHP"),
		"sass":   cStyle("Sass"),
		"swift":  cStyle("Swift"),
		"cfc":    cStyle("ColdFusionScript"),
		"css":    cStyle("CSS"),
		"cu":     cStyle("CUDA"),
		"cuh":    cStyle("CUDA Header"),
		"d":      cStyle("D"),
		"dart":   cStyle("Dart"),
		"dts":    cStyle("DeviceTree"),
		"vert":   cStyle("GLSL"),
		"jai":    cStyle("Jai"),
		"jsx":    cStyle("Jsx"),
		"less":   cStyle("Less"),
		"lds":    cStyle("LinkerScript"),
		"qcl":    cStyle("Qcl"),
		"qml":    cStyle("QML"),
		"rs":     cStyle("Rust"),
		"scala":  cStyle("Scala"),
		"styl":   cStyle("Stylus"),
		"ts":     cStyle("TypeScript"),
		"tsx":    cStyle("Typescript JSX"),
		"uc":     cStyle("UnrealScript"),
		"y":      cStyle("Yacc"),

		"txt":  noComments("Plain text"),
		"hex":  noComments("Hex"),
		"ihex": noComments("Intel Hex"),
		"rst":  noComments("reStructuredText"),

		"html":   htmlStyle("HTML"),
		"xml":    htmlStyle("XML"),
		"polly":  htmlStyle("Polly"),
		"cshtml": htmlStyle("Razor"),
		"rhtml":  htmlStyle("RubyHtml"),

		"makefile": shStyle("Makefile"),
		"r":        shStyle("R"),
		"yml":      shStyle("YAML"),
		"in":       shStyle("Autoconf"),
		"awk":      shStyle("Awk"),
		"sh":       shStyle("Bourne Shell"),
		"csh":      shStyle("C Shell"),
		"nim":      shStyle("Nim"),
		"sls":      shStyle("SaltStack"),
		"tcl":      shStyle("Tcl"),
		"toml":     shStyle("Toml"),

		"v":   mlStyle("Coq"),
		"ml":  mlStyle("OCaml"),
		"sml": mlStyle("SML"),
		"wl":  mlStyle("Wolfram"),

		"oz": prologStyle("Oz"),
		"p":  prologStyle("Prolog"),

		"md":       language("Markdown", singleComment("[//]: # "), xmlComment),
		"pas":      language("Pascal", models.Comment{"{", "}"}, models.Comment{"{*", "*}"}),
		"pl":       language("Perl", shSComment, models.Comment{"=begin", "=cut"}),
		"py":       language("Python", shSComment, models.Comment{"\"\"\"", "\"\"\""}),
		"rb":       language("Ruby", shSComment, models.Comment{"=begin", "=end"}),
		"sql":      language("SQL", sqlSComment, clangMComment),
		"ads":      language("Ada", sqlSComment, noneComment),
		"agda":     language("Agda", sqlSComment, agdaMComment),
		"asp":      language("ASP", aspSComment, xmlComment),
		"aspx":     language("ASP.NET", aspSComment, xmlComment),
		"asm":      language("Assembly", shSComment, clangMComment),
		"bat":      language("Batch", singleComment("REM"), noneComment),
		"coffee":   language("CoffeeScript", shSComment, models.Comment{"###", "###"}),
		"fs":       language("F#", clangSComment, mlMComment),
		"hs":       language("Haskell", sqlSComment, agdaMComment),
		"erl":      language("Erlang", singleComment("%"), noneComment),
		"forth":    language("Forth", singleComment("\\"), models.Comment{"(", ")"}),
		"f":        language("FORTRAN Legacy", singleComment("!"), noneComment),
		"f03":      language("FORTRAN Modern", singleComment("!"), noneComment),
		"hbs":      language("Handlebars", xmlComment, models.Comment{"{{!", "}}"}),
		"idr":      language("Idris", sqlSComment, agdaMComment),
		"ini":      language("INI", singleComment(";"), noneComment),
		"jl":       language("Julia", shSComment, models.Comment{"#=", "=#"}),
		"lean":     language("Lean", sqlSComment, models.Comment{"/-", "-/"}),
		"lisp":     language("Lisp", singleComment(";"), lispMComment),
		"lua":      language("Lua", sqlSComment, models.Comment{"--[[", "]]"}),
		"mustache": language("Mustache", noneComment, models.Comment{"{{!", "}}"}),
		"nix":      language("Nix", shSComment, clangMComment),
		"proto":    language("Protobuf", clangSComment, noneComment),
		"zig":      language("Zig", clangSComment, noneComment),
		"arr":      language("Pyret", shSComment, lispMComment),
		"tf":       language("Terraform", shSComment, clangMComment),
		"tex":      language("TeX", singleComment("%"), noneComment),
	}
)
