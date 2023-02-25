package main

import (
	"fmt"

	abstractfactory "github.com/IvanProg00/go-patterns/abstract_factory"
	"github.com/IvanProg00/go-patterns/adapter"
	"github.com/IvanProg00/go-patterns/bridge"
	"github.com/IvanProg00/go-patterns/builder"
	"github.com/IvanProg00/go-patterns/chainofresponsibility"
	"github.com/IvanProg00/go-patterns/composite"
	"github.com/IvanProg00/go-patterns/decorator"
	"github.com/IvanProg00/go-patterns/facade"
	factorymethod "github.com/IvanProg00/go-patterns/factory_method"
	"github.com/IvanProg00/go-patterns/iterator"
	"github.com/IvanProg00/go-patterns/observer"
	"github.com/IvanProg00/go-patterns/singleton"
	"github.com/IvanProg00/go-patterns/snapshot"
	"github.com/IvanProg00/go-patterns/state"
	"github.com/IvanProg00/go-patterns/strategy"
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
	printLine()

	fmt.Println("Composite")
	composite.Run()
	printLine()

	fmt.Println("Chain of Responsibility")
	chainofresponsibility.Run()
	printLine()

	fmt.Println("Singleton")
	singleton.Run()
	printLine()

	fmt.Println("Strategy")
	strategy.Run()
	printLine()

	fmt.Println("Iterator")
	iterator.Run()
	printLine()

	fmt.Println("Snapshot")
	snapshot.Run()
}

func printLine() {
	fmt.Println("--------------------")
}
