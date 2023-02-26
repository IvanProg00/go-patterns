package vehicle

import "fmt"

type Passenger struct {
	Dispatcher
}

func (p *Passenger) Arrive() {
	if !p.CanArrive(p) {
		fmt.Println("passengers: departure delayed")

		return
	}

	fmt.Println("passengers: take your seats")
}

func (p *Passenger) Go() {
	fmt.Println("passengers: departure")
	p.Dispatcher.NotifyAboutGo()
}

func (p *Passenger) PermitArrive() {
	fmt.Println("passengers: take your seats")
	p.Arrive()
}
