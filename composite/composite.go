package composite

import "github.com/IvanProg00/go-patterns/composite/computer"

func Run() {
	cpu1 := &computer.CPU{
		Name:        "CPU-1",
		Description: "Processor 1",
	}
	cpu2 := &computer.CPU{
		Name:        "CPU-2",
		Description: "Processor 2",
	}
	gpu1 := &computer.GPU{
		Name:        "Radeon",
		Description: "Video card 1",
	}
	gpu2 := &computer.GPU{
		Name:        "GeForce",
		Description: "Video card 2",
	}

	motherboard := &computer.Motherboard{
		Name:        "Gigabyte",
		Description: "Mother card 1",
		Components: []computer.Component{
			cpu1, cpu2,
			gpu1, gpu2,
		},
	}

	pc := &computer.PC{
		Name:        "PC",
		Description: "Personal Computer",
		Components: []computer.Component{
			motherboard,
		},
	}

	pc.Search("PC")
	pc.Search("Gigabyte")
	pc.Search("CPU-2")
	pc.Search("Radeon")
}
