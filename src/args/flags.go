package args

import "flag"

var (
	dirPtr     = flag.String(dirFlag, "", "count lines of all files in directory")
	filePtr    = flag.String(fileFlag, "", "count lines of one file")
	filesPtr   = flag.String(filesFlag, "", "count lines of multiple files")
	jsonOutPtr = flag.Bool("j", false, "generate json result (default is std output)")
	xmlOutPtr  = flag.Bool("x", false, "generate xml result (default is std output)")
	ymlOutPtr  = flag.Bool("y", false, "generate yaml result (default is std output)")
	OutPathPtr = flag.String("c", "./", "path to result file(s)")
	help       = flag.Bool("help", false, "read usage info")
	h          = flag.Bool("h", false, "read usage info")
)
