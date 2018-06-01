package main

import (
	"fmt"
	"github.com/YuriyLisovskiy/sloc/src/out"
	"github.com/YuriyLisovskiy/sloc/src/args"
	"github.com/YuriyLisovskiy/sloc/src/parser"
	"github.com/YuriyLisovskiy/sloc/src/models"
)

func main() {
	dir, file, multipleFiles, jsonOut, xmlOut, ymlOut, help, err := args.Parse()
	if err == nil {
		if !help {
			var result []models.Lang
			var total *models.Lang
			if dir != "" {
				res, newTotal := parser.ParseDirectory(dir, nil)
				result = res
				total = &newTotal
			} else if file != "" {
				resFile, err := parser.ParseFile(file)
				if err == nil {
					result = []models.Lang{resFile}
				}
			} else if len(multipleFiles) > 0 {
				res, newTotal := parser.ParseMultiple(multipleFiles)
				result = res
				total = &newTotal
			}
			if jsonOut {
				if err = out.ToJson(result, total); err != nil {
					fmt.Println(err.Error())
				}
			}
			if xmlOut {
				if out.ToXml(result, total); err != nil {
					fmt.Println(err.Error())
				}
			}
			if ymlOut {
				if out.ToYaml(result, total); err != nil {
					fmt.Println(err.Error())
				}
			}
			if out.ToStd(result, total); err != nil {
				fmt.Println(err.Error())
			}
		}
	} else {
		println(err.Error())
	}
}
