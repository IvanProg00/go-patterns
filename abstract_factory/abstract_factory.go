package abstractfactory

import (
	"fmt"
	"log"

	"github.com/IvanProg00/go-patterns/abstract_factory/factory"
)

func Run() {
	fmt.Println("DELL")

	dellFactory, err := factory.GetComputerFactory(factory.DellBrand)
	if err != nil {
		log.Fatalf("get dell factory: %v", err)
	}

	dellMonitor := dellFactory.GetMonitor()
	printMonitor(dellMonitor)

	dellComputer := dellFactory.GetComputer()
	printComputer(dellComputer)

	fmt.Println("HP")

	hpFactory, err := factory.GetComputerFactory(factory.HPBrand)
	if err != nil {
		log.Fatalf("get hp factory: %v", err)
	}

	hpMonitor := hpFactory.GetMonitor()
	printMonitor(hpMonitor)

	hpComputer := hpFactory.GetComputer()
	printComputer(hpComputer)

	_, err = factory.GetComputerFactory("not existing")
	log.Println(err)
}

func printMonitor(m factory.IMonitor) {
	fmt.Printf("monitor: size[%d]\n", m.GetSize())
}

func printComputer(c factory.IComputer) {
	fmt.Printf("computer: cpu[%d] memory[%d]\n", c.GetCPU(), c.GetMemory())
}
