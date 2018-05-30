package args

import "flag"

var (
	dirFlag    = "d"
	fileFlag   = "f"
	filesFlag  = "mf"
	dirPtr     = flag.String(dirFlag, "", "count lines of all files in directory")
	filePtr    = flag.String(fileFlag, "", "count lines of one file")
	filesPtr   = flag.String(filesFlag, "", "count lines of multiple files")
	jsonOutPtr = flag.Bool("json", false, "generate json result (default is std output)")
)
