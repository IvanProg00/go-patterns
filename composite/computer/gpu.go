package computer

import "fmt"

type GPU struct {
	Name        string
	Description string
}

func (g *GPU) Search(name string) {
	if g.Name == name {
		fmt.Printf("Component %q found: %q\n", g.Name, g.Description)
	}
}

func (g *GPU) GetName() string {
	return g.Name
}
