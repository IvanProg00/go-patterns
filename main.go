package main

import (
	"fmt"

	abstractfactory "github.com/IvanProg00/go-patterns/abstract_factory"
	"github.com/IvanProg00/go-patterns/adapter"
	"github.com/IvanProg00/go-patterns/bridge"
	"github.com/IvanProg00/go-patterns/builder"
	"github.com/IvanProg00/go-patterns/decorator"
	"github.com/IvanProg00/go-patterns/facade"
	factorymethod "github.com/IvanProg00/go-patterns/factory_method"
	"github.com/IvanProg00/go-patterns/observer"
	"github.com/IvanProg00/go-patterns/state"
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
	printLine()

	fmt.Println("Facade")
	facade.Run()
	printLine()

	fmt.Println("Decorator")
	decorator.Run()
	printLine()

	fmt.Println("State")
	state.Run()
	printLine()

	fmt.Println("Observer")
	observer.Run()
}

func printLine() {
	fmt.Println("--------------------")
}
