package parser

var normalizeLangTestData = []struct {
	input    string
	expected string
}{
	{"cs", "cs"},
	{"h", "c"},
	{"hpp", "cpp"},
	{"yaml", "yml"},
	{"htm", "html"},
}

var getExtTestData = []struct {
	input    string
	expected string
}{
	{"file.cs", "cs"},
	{"file.hpp", "hpp"},
	{"file.go", "go"},
}
