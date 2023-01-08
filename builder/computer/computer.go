package computer

import "fmt"

type Computer struct {
	Brand       string
	Core        int
	Memory      int
	GraphicCard int
}

func (c *Computer) Print() {
	fmt.Printf("%s core: [%d] mem: [%d] graphic card[%d]\n", c.Brand, c.Core, c.Memory, c.GraphicCard)
}
