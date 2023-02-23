package computer

import "fmt"

type PC struct {
	Name        string
	Description string
	Components  []Component
}

func (pc *PC) Search(name string) {
	if pc.Name == name {
		fmt.Printf("Component %q found: %q\n", pc.Name, pc.Description)
		return
	}

	for _, comp := range pc.Components {
		comp.Search(name)
	}
}

func (pc *PC) GetName() string {
	return pc.Name
}
