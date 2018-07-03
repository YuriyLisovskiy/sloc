package cli

var (
	err       = "sloc args error: '%s'"
	argsError = "sloc args error: use %s or %s"
)

const (
	version     = "sloc 0.0.1"
	description = "A utility for counting source lines of code"
	author      = "Yuriy Lisovskiy (c) 2018 <https://github.com/YuriyLisovskiy>"
	usage       = "USAGE:\n" +
		"    sloc [COMMANDS] [FLAGS] [OPTIONS]\n\n" +
		"COMMANDS:\n" +
		"    count\tCounts lines using FLAGS\n" +
		"    help\tPrints help information\n" +
		"    version\tPrints version information\n\n" +
		"FLAGS:\n" +
		"    -f, --file\t\tSets file to count lines\n" +
		"    -d, --directory\tSets folder to count lines\n" +
		"    -m, --multiple\tSets files and(or) folders to count lines using \"\"\n\n" +
		"OPTIONS:\n" +
		"    -e, --exclude\tSets file(s) and(or) folder(s) to exclude from counting lines using \"\"\n" +
		"    -j, --json\t\tWrites result to json file\n" +
		"    -x, --xml\t\tWrites result to xml file\n" +
		"    -y, --yaml\t\tWrites result to yaml file\n" +
		"    -l, --log\t\tPrints log\n"
)
