package cli

import "flag"

var (
	countCmd   = flag.NewFlagSet("count", flag.ExitOnError)
	fPtr       = countCmd.String("f", "", "Sets file to count lines")
	filePtr    = countCmd.String("file", "", "Sets file to count lines")
	dPtr       = countCmd.String("d", "", "Sets folder to count lines")
	dirPtr     = countCmd.String("directory", "", "Sets folder to count lines")
	mPtr       = countCmd.String("m", "", "Sets files and(or) folders to count lines using \"\"")
	multPtr    = countCmd.String("multiple", "", "Sets files and(or) folders to count lines using \"\"")
	ePtr       = countCmd.String("e", "", "Sets file(s) and(or) folder(s) to exclude from counting lines")
	excludePtr = countCmd.String("exclude", "", "Sets file(s) and(or) folder(s) to exclude from counting lines")
	jsonOut    = countCmd.Bool("json", false, "Writes result to json file")
	jOut       = countCmd.Bool("j", false, "Writes result to json file")
	xmlOut     = countCmd.Bool("xml", false, "Writes result to xml file")
	xOut       = countCmd.Bool("x", false, "Writes result to xml file")
	yamlOut    = countCmd.Bool("yaml", false, "Writes result to yaml file")
	yOut       = countCmd.Bool("y", false, "Writes result to yaml file")
	printLog   = countCmd.Bool("log", false, "Prints log")
	printL     = countCmd.Bool("l", false, "Prints log")
)
