package cli

import (
	"os"
	"log"
	"fmt"
	"github.com/YuriyLisovskiy/sloc/src/out"
	"github.com/YuriyLisovskiy/sloc/src/models"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

type CLI struct{}

func (cli *CLI) printAbout() {
	print(fmt.Sprintf("%s\n%s\n%s\n\n%s\n", version, author, description, usage))
}

func (cli *CLI) normalizeFlags() {
	if len(*fPtr) == 0 {
		fPtr = filePtr
	}
	if len(*dPtr) == 0 {
		dPtr = dirPtr
	}
	if len(*mPtr) == 0 {
		mPtr = multPtr
	}
	if len(*ePtr) == 0 {
		ePtr = excludePtr
	}
	if !*printL {
		printL = printLog
	}
	if !*jsonOut {
		jsonOut = jOut
	}
	if !*xmlOut {
		xmlOut = xOut
	}
	if !*yamlOut {
		yamlOut = yOut
	}
}

func (cli *CLI) Run() {
	if len(os.Args) < 2 {
		cli.printAbout()
		return
	}
	switch os.Args[1] {
	case "count":
		err := countCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "help":
		println(usage)
	case "version":
		println(version)
	default:
		cli.printAbout()
	}
	if countCmd.Parsed() {
		cli.normalizeFlags()
		dir, file, multiple, err := Parse(*fPtr, *dPtr, *mPtr, *ePtr)
		if err == nil {
			var result []models.Lang
			var total *models.Lang
			if dir != "" {
				res, newTotal := parser.ParseDirectory(dir, nil, *printL)
				result = res
				total = &newTotal
			} else if file != "" {
				resFile, err := parser.ParseFile(file, *printL)
				if err == nil {
					result = []models.Lang{resFile}
				}
			} else if len(multiple) > 0 {
				res, newTotal := parser.ParseMultiple(multiple, *printL)
				result = res
				total = &newTotal
			}
			if *jsonOut {
				if err = out.ToJson(result, total); err != nil {
					fmt.Println(err.Error())
				}
			}
			if *xmlOut {
				if out.ToXml(result, total); err != nil {
					fmt.Println(err.Error())
				}
			}
			if *yamlOut {
				if out.ToYaml(result, total); err != nil {
					fmt.Println(err.Error())
				}
			}
			if out.ToStd(result, total); err != nil {
				fmt.Println(err.Error())
			}
		} else {
			println(err.Error())
		}
	}
}
