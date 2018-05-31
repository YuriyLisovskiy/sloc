package utils

import (
	"testing"
	"github.com/YuriyLisovskiy/sloc/src/utils"
)

func TestGetTemplates(test *testing.T) {
	for _, td := range getTemplatesTestData {
		actualLine, actualHeader, actualFormat := utils.GetTemplates(td.langs)
		if actualLine != td.expectedLine {
			test.Errorf("utils.TestGetTemplates: expected %s, actual %s", td.expectedLine, actualLine)
		}
		if actualHeader != td.expectedHeader {
			test.Errorf("utils.TestGetTemplates: expected %s, actual %s", td.expectedHeader, actualHeader)
		}
		if actualFormat != td.expectedFormat {
			test.Errorf("utils.TestGetTemplates: expected %s, actual %s", td.expectedFormat, actualFormat)
		}
	}
}

func TestGetFormatTemplate(test *testing.T) {
	for _, td := range getFormatTemplateTestData {
		actual := utils.GetFormatTemplate(td.indent, td.dataType)
		if actual != td.expected {
			test.Errorf("utils.TestGetFormatTemplate: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestItos(test *testing.T) {
	for _, td := range itosTestData {
		actual := utils.Itos(td.input)
		if actual != td.expected {
			test.Errorf("utils.TestItos: expected %s, actual %s", td.expected, actual)
		}
	}
}

func TestItol(test *testing.T) {
	for _, td := range itolTestData {
		actual := utils.Itol(td.input)
		if actual != td.expected {
			test.Errorf("utils.TestItol: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestFindMax(test *testing.T) {
	for _, td := range findMaxTestData {
		actual := utils.FindMax(td.first, td.second)
		if actual != td.expected {
			test.Errorf("utils.TestFindMax: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestFindFieldWithMaxLen(test *testing.T) {
	for _, td := range findFieldWithMaxLenTestData {
		actual := utils.FindFieldWithMaxLen(td.input)
		if actual != td.expected {
			test.Errorf("utils.TestFindFieldWithMaxLen: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestFindIndentLen(test *testing.T) {
	for _, td := range findIndentLenTestData {
		actual := utils.FindIndentLen(td.input)
		if actual != td.expected {
			test.Errorf("utils.TestFindIndentLen: expected %d, actual %d", td.expected, actual)
		}
	}
}

func TestAppendLangData(test *testing.T) {
	for _, td := range appendLangDataTestData {
		actual := utils.AppendLangData(td.lang, td.format)
		if actual != td.expected {
			test.Errorf("utils.TestAppendLangData: expected %s, actual %s", td.expected, actual)
		}
	}
}
