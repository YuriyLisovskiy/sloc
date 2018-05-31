package main

import (
	"github.com/YuriyLisovskiy/sloc/src/args"
	"github.com/YuriyLisovskiy/sloc/src/utils"
	"github.com/YuriyLisovskiy/sloc/src/parser"
	"fmt"
)

func main() {
	dir, file, multipleFiles, jsonOut, xmlOut, ymlOut, help, err := args.Parse()
	if err == nil {
		if !help {
			var res []parser.Lang
			var total parser.Lang
			if dir != "" {
				res, total = parser.Parse(parser.ReadDir(dir))
			} else if file != "" {
				res, total = parser.Parse([]string{file})
			} else if len(multipleFiles) > 0 {
				res, total = parser.Parse(multipleFiles)
			}
			if jsonOut {
				if err = utils.OutputToJson(res, total); err != nil {
					fmt.Println(err.Error())
				}

			}
			if xmlOut {
				if utils.OutputToXml(res, total); err != nil {
					fmt.Println(err.Error())
				}
			}
			if ymlOut {
				if utils.OutputToYaml(res, total); err != nil {
					fmt.Println(err.Error())
				}
			}
			if utils.OutputToStd(res, total); err != nil {
				fmt.Println(err.Error())
			}
		}
	} else {
		println(err.Error())
	}
}
