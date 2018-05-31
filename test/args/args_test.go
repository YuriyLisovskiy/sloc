package args

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/args"
)

func TestSplitMultFiles(test *testing.T) {
	for _, td := range SplitMultFilesTestData {
		actual := args.SplitMultFiles(td.input)
		for i, a := range actual {
			if a != td.expected[i] {
				test.Errorf("args.TestSplitMultFiles: expected %s, actual %s", td.expected, actual)
			}
		}
	}
}

func TestValidateParams(test *testing.T) {
	for _, td := range ValidateTestData {
		actual := args.ValidateParams(td.dir, td.file, td.files) == nil
		if actual != td.expected {
			test.Errorf("args.TestValidate: expected %t, actual %t", td.expected, actual)
		}
	}
}
