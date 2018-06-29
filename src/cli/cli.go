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

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  count")
	fmt.Println("    -f\n\tset file to count lines")
	fmt.Println("    -d\n\tset folder to count lines")
	fmt.Println("    -m\n\tset set files and(or) folders to count lines using \"\"")
	fmt.Println("    -e\n\tset file(s) and(or) folder(s) to exclude from counting lines using \"\"")
	fmt.Println("    -json\n\twrite result to json file")
	fmt.Println("    -xml\n\twrite result to xml file")
	fmt.Print("    -yaml\n\twrite result to yaml file\n\n")
	fmt.Print("  help\n\tread usage\n\n")
}

func (cli *CLI) validateArgs() bool {
	if len(os.Args) < 2 {
		cli.printUsage()
		return false
	}
	return true
}

func (cli *CLI) Run() {
	if !cli.validateArgs() {
		return
	}
	f := countCmd.String("f", "", "set file to count lines")
	d := countCmd.String("d", "", "set folder to count lines")
	m := countCmd.String("m", "", "set files and(or) folders to count lines using \"\"")
	e := countCmd.String("e", "", "set file(s) and(or) folder(s) to exclude from counting lines")
	jsonOut := countCmd.Bool("json", false, "write result to json file")
	xmlOut := countCmd.Bool("xml", false, "write result to xml file")
	yamlOut := countCmd.Bool("yaml", false, "write result to yaml file")
	switch os.Args[1] {
	case "count":
		err := countCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "help":
		cli.printUsage()
		return
	default:
		cli.printUsage()
		return
	}
	if countCmd.Parsed() {
		dir, file, multiple, err := Parse(*f, *d, *m, *e)
		if err == nil {
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
			} else if len(multiple) > 0 {
				res, newTotal := parser.ParseMultiple(multiple)
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
	} else {
		panic("Error")
	}
}
