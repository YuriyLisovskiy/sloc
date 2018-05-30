package utils

import "strings"

var (
	stdOutLine = strings.Join(make([]string, 61), "-") + "\n"
	stdOutHeader = stdOutLine + "Language\tFiles\tLines\tBlank\tComments\tCode\n" + stdOutLine
	stdOutFormatData = "%s\t\t%d\t%d\t%d\t%d\t\t%d\n"
)
