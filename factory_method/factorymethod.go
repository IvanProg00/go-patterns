package factorymethod

import (
	"fmt"
	"log"

	"github.com/IvanProg00/go-patterns/factory_method/computer"
)

func Run() {
	laptop, err := computer.New(computer.LaptopType)
	if err != nil {
		log.Fatalf("new laptop: %v", err)
	}
	printComputer(laptop)

	server, err := computer.New(computer.ServerType)
	if err != nil {
		log.Fatalf("new server: %v", err)
	}
	printComputer(server)

	pc, err := computer.New(computer.PersonalComputerType)
	if err != nil {
		log.Fatalf("new personal computer: %v", err)
	}
	printComputer(pc)

	_, err = computer.New("undefined computer")
	if err != nil {
		fmt.Printf("ERROR: new undefined computer: %v", err)
	}
}

func printComputer(comp computer.Computer) {
	fmt.Printf("type: [%s]; cpu: [%d]; os: [%s]\n", comp.GetType(), comp.GetCPU(), comp.GetOS())
}
