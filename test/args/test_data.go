package args

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
		dir string
		file string
		files []string
		expected bool
	} {
		{"dir/", "file.ext", nil, false},
		{"dir/", "", []string{"file1.ext1", "file2.ext2"}, false},
		{"", "file.ext", []string{"file1.ext1", "file2.ext2"}, false},
		{"dir/", "", nil, true},
		{"", "file.ext", nil, true},
		{"", "", []string{"file1.ext1", "file2.ext2"}, true},
	}
)
