package main

import (
	"github.com/YuriyLisovskiy/sloc/src/parser"
	"github.com/YuriyLisovskiy/sloc/src/utils"
)

func main() {
/*
	d, f, ff, j, err := args.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(d, f, ff, j)
*/
	path := "../../src/parser/"
	languages, total := parser.Parse([]string{path + "lang.go"})
	utils.OutputToStd(languages, total)
}
