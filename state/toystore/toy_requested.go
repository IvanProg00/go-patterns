package toystore

import "fmt"

type ToyRequestedState struct {
	toyStore *ToyStore
}

func (s *ToyRequestedState) AddToys(count int) error {
	return ErrToyDispenseInProgress
}

func (s *ToyRequestedState) RequestToy() error {
	return ErrToyAlreadyRequested
}

func (s *ToyRequestedState) GiveMoney(money int) error {
	if money < s.toyStore.toysPrice {
		return ErrReceivedMoneyIsLess
	}

	fmt.Println("money received")

	s.toyStore.SetState(s.toyStore.hasMoney)

	return nil
}

func (s *ToyRequestedState) DispenseToy() error {
	return ErrGiveMoneyFirst
}
