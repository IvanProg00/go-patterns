package computer

import (
	"fmt"
)

type CPU struct {
	Name        string
	Description string
}

func (c *CPU) Search(name string) {
	if c.Name == name {
		fmt.Printf("Component %q found: %q\n", c.Name, c.Description)
	}
}

func (c *CPU) GetName() string {
	return c.Name
}
