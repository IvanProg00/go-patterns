package computer

import "fmt"

type Motherboard struct {
	Name        string
	Description string
	Components  []Component
}

func (mb *Motherboard) Search(name string) {
	if mb.Name == name {
		fmt.Printf("Component %q found: %q\n", mb.Name, mb.Description)
		return
	}

	for _, comp := range mb.Components {
		comp.Search(name)
	}
}

func (mb *Motherboard) GetName() string {
	return mb.Name
}
