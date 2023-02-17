package toystore

import "fmt"

type HasMoneyState struct {
	ToyStore *ToyStore
}

func (s *HasMoneyState) AddToys(count int) error {
	return ErrToyDispenseInProgress
}

func (s *HasMoneyState) RequestToy() error {
	return ErrToyDispenseInProgress
}

func (s *HasMoneyState) GiveMoney(money int) error {
	return ErrNoToys
}

func (s *HasMoneyState) DispenseToy() error {
	fmt.Println("dispense toy")
	s.ToyStore.toysCount--

	if s.ToyStore.toysCount == 0 {
		s.ToyStore.SetState(s.ToyStore.noToys)
	} else {
		s.ToyStore.SetState(s.ToyStore.hasToy)
	}

	return nil
}
