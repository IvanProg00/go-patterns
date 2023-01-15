package main

import (
	"fmt"

	abstractfactory "github.com/IvanProg00/go-patterns/abstract_factory"
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
	printLine()
	abstractfactory.Run()
}

func printLine() {
	fmt.Println("--------------------")
}
