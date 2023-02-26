package vehicle

import "fmt"

type Cargo struct {
	Dispatcher
}

func (c *Cargo) Arrive() {
	if !c.CanArrive(c) {
		fmt.Println("truck: loading is completed, passengers are on the platform")
		return
	}

	fmt.Println("truck: shipped")
}

func (c *Cargo) Go() {
	fmt.Println("truck: successfully loaded")
	c.Dispatcher.NotifyAboutGo()
}

func (c *Cargo) PermitArrive() {
	fmt.Println("truck: loading cargo")
	c.Arrive()
}
