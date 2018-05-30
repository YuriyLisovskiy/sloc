package main

import (
	"github.com/YuriyLisovskiy/sloc/src/args"
	"github.com/YuriyLisovskiy/sloc/src/parser"
	"github.com/YuriyLisovskiy/sloc/src/utils"
)

func main() {
	dir, _, _, _, err := args.Parse()
	if err != nil {
		panic(err)
	}
	res, other, total := parser.Parse(parser.ReadDir(dir))
	utils.OutputToStd(res, total, &other)
}
