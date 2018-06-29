package cli

import (
	"os"
	"log"
	"fmt"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  count")
	fmt.Println("    -f\n\tset file(s) to count lines")
	fmt.Println("    -d\n\tset folder(s) to count lines")
	fmt.Println("    -e\n\tset file(s) and(or) folder(s) to exclude from counting lines")
	fmt.Println("    -json\n\twrite result to json file")
	fmt.Println("    -xml\n\twrite result to xml file")
	fmt.Print("    -yaml\n\twrite result to yaml file\n\n")
	fmt.Print("  help\n\tRead usage\n\n")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

//	f := countCmd.String("f", "", "set file(s) to count lines")
//	d := countCmd.String("d", "", "set folder(s) to count lines")
//	e := countCmd.String("e", "", "set file(s) and(or) folder(s) to exclude from counting lines")
//	j := countCmd.Bool("json", false, "write result to json file")
//	x := countCmd.Bool("xml", false, "write result to xml file")
//	y := countCmd.Bool("yaml", false, "write result to yaml file")

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

	}
}
