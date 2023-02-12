package decorator

import (
	"fmt"

	"github.com/IvanProg00/go-patterns/decorator/pc"
)

func Run() {
	base := pc.BasePC{}

	home := pc.HomePC{
		CPU:          4,
		GraphicsCard: 1,
		Wrapper:      base,
	}

	server := pc.ServerPC{
		CPU:     16,
		Memory:  256,
		Wrapper: base,
	}

	fmt.Printf("Base: price=%f\n", base.GetPrice())
	fmt.Printf("Home: cpu=%d, cards=%d, price=%f\n", home.CPU, home.GraphicsCard, home.GetPrice())
	fmt.Printf("Server: cpu=%d, memory=%d, price=%f\n", server.CPU, server.Memory, server.GetPrice())
}
