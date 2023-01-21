package main

import (
	"fmt"

	abstractfactory "github.com/IvanProg00/go-patterns/abstract_factory"
	"github.com/IvanProg00/go-patterns/adapter"
	"github.com/IvanProg00/go-patterns/bridge"
	"github.com/IvanProg00/go-patterns/builder"
	factorymethod "github.com/IvanProg00/go-patterns/factory_method"
)

func main() {
	fmt.Println("Adapter")
	adapter.Run()
	printLine()

	fmt.Println("Factory method")
	factorymethod.Run()
	printLine()

	fmt.Println("Builder")
	builder.Run()
	printLine()

	fmt.Println("Abstract factory")
	abstractfactory.Run()
	printLine()

	fmt.Println("Bridge")
	bridge.Run()
}

func printLine() {
	fmt.Println("--------------------")
}
