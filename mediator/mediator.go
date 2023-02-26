package mediator

import "github.com/IvanProg00/go-patterns/mediator/vehicle"

func Run() {
	stationManager := vehicle.NewStationManager()

	bus := &vehicle.Passenger{
		Dispatcher: stationManager,
	}

	cargo := &vehicle.Cargo{
		Dispatcher: stationManager,
	}

	bus.Arrive()
	cargo.Arrive()
	bus.Go()
}
