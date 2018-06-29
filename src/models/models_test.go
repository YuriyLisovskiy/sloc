package models

import "testing"

var toStringTestData = []struct {
	lang     Lang
	expected string
}{
	{
		Lang{"Language1", 100, 2322, 233, 33, 33},
		"Language1 100 2322 233 33 33",
	},
	{
		Lang{"Java", 1002, 222122, 23223, 3483, 3883},
		"Java 1002 222122 23223 3483 3883",
	},
	{
		Lang{"Python", 10, 322, 263, 2, 98},
		"Python 10 322 263 2 98",
	},
}

func TestToString(test *testing.T) {
	for _, td := range toStringTestData {
		actual := td.lang.ToString()
		if actual != td.expected {
			test.Errorf("models.TestToString: expected %s, actual %s", td.expected, actual)
		}
	}
}
