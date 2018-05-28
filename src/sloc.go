package main

import (
	"fmt"
	"github.com/YuriyLisovskiy/sloc/src/args"
)

func main() {
	println("Hello, sloc!")
	d, f, ff, j, err := args.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(d, f, ff, j)
}
