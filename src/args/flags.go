package args

import "flag"

var (
	h          = flag.Bool("h", false, "read usage info")
	help       = flag.Bool("help", false, "read usage info")
	filePtr    = flag.String(fileFlag, "", "count lines of one file")
	OutPathPtr = flag.String("c", "./", "path to result file(s)")
	filesPtr   = flag.String(filesFlag, "", "count lines of multiple files")
	dirPtr     = flag.String(dirFlag, "", "count lines of all files in directory")
	excludePtr = flag.String("e", "", "exclude file(s) and(or) directory(s)")
	xmlOutPtr  = flag.Bool("x", false, "generate xml result (default is std output)")
	jsonOutPtr = flag.Bool("j", false, "generate json result (default is std output)")
	ymlOutPtr  = flag.Bool("y", false, "generate yaml result (default is std output)")
)
