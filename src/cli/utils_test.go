package cli

import "testing"

var (
	SplitMultFilesTestData = []struct {
		input    string
		expected []string
	}{
		{"file1.cs file2.go file3.cpp", []string{"file1.cs", "file2.go", "file3.cpp"}},
		{"file1.cs", []string{"file1.cs"}},
		{" file1.cs    file2.hpp  ", []string{"file1.cs", "file2.hpp"}},
	}
	ValidateTestData = []struct{
		dirs []string
		files []string
		expected bool
	} {
		{[]string{"dir/"}, []string{"file.ext"}, false},
		{[]string{"dir/"}, nil, true},
		{nil, []string{"file.ext"}, true},
		{nil, nil, false},
	}
)

func TestSplitMultFiles(test *testing.T) {
	for _, td := range SplitMultFilesTestData {
		actual := splitMultFiles(td.input)
		for i, a := range actual {
			if a != td.expected[i] {
				test.Errorf("args.TestSplitMultFiles: expected %s, actual %s", td.expected, actual)
			}
		}
	}
}

func TestValidateParams(test *testing.T) {
	for _, td := range ValidateTestData {
		actual := ValidateParams(td.dirs, td.files) == nil
		if actual != td.expected {
			test.Errorf("args.TestValidate: expected %t, actual %t", td.expected, actual)
		}
	}
}

