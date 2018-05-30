package main

import (
	"github.com/YuriyLisovskiy/sloc/src/args"
	"github.com/YuriyLisovskiy/sloc/src/utils"
	"github.com/YuriyLisovskiy/sloc/src/parser"
)

func main() {
	dir, _, _, _, err := args.Parse()
	if err != nil {
		panic(err)
	}
	res, other, total := parser.Parse(parser.ReadDir(dir))
	utils.OutputToStd(res, total, &other)
}
