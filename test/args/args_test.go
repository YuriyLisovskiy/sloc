package args

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/args"
)

var SplitMultFilesTestData = []struct {
	input    string
	expected []string
}{
	{"file1.cs file2.go file3.cpp", []string{"file1.cs", "file2.go", "file3.cpp"}},
	{"file1.cs", []string{"file1.cs"}},
	{" file1.cs    file2.hpp  ", []string{"file1.cs", "file2.hpp"}},
}

func TestSplitMultFiles(test *testing.T) {
	for _, td := range SplitMultFilesTestData {
		actual := args.SplitMultFiles(td.input)
		for i, a := range actual {
			if a != td.expected[i] {
				test.Errorf("args.SplitMultFiles(%s): expected %s, actual %s", td.input, td.expected, actual)
			}
		}
	}
}
