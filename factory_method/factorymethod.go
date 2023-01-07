package factorymethod

import (
	"fmt"
	"log"

	"github.com/IvanProg00/go-patterns/factory_method/computer"
)

func Run() {
	laptop, err := computer.Get(computer.LaptopType)
	if err != nil {
		log.Fatalf("new laptop: %v", err)
	}
	printComputer(laptop)

	server, err := computer.Get(computer.ServerType)
	if err != nil {
		log.Fatalf("new server: %v", err)
	}
	printComputer(server)

	pc, err := computer.Get(computer.PersonalComputerType)
	if err != nil {
		log.Fatalf("new personal computer: %v", err)
	}
	printComputer(pc)

	_, err = computer.Get("undefined computer")
	if err != nil {
		fmt.Printf("ERROR: new undefined computer: %v\n", err)
	}
}

func printComputer(comp computer.Computer) {
	fmt.Printf("type: [%s]; cpu: [%d]; os: [%s]\n", comp.GetType(), comp.GetCPU(), comp.GetOS())
}
