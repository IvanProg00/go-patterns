package main

import (
	"fmt"

	"github.com/IvanProg00/go-patterns/adapter"
	"github.com/IvanProg00/go-patterns/builder"
	factorymethod "github.com/IvanProg00/go-patterns/factory_method"
)

func main() {
	adapter.Run()
	printLine()
	factorymethod.Run()
	printLine()
	builder.Run()
}

func printLine() {
	fmt.Println("--------------------")
}
