package toystore

import (
	"fmt"
)

type HasToyState struct {
	toyStore *ToyStore
}

func (s *HasToyState) AddToys(count int) error {
	fmt.Printf("toys added %d\n", count)
	s.toyStore.IncrementItemCount(count)

	return nil
}

func (s *HasToyState) RequestToy() error {
	if s.toyStore.toysCount == 0 {
		s.toyStore.SetState(s.toyStore.noToys)
		return ErrNoToys
	}

	fmt.Println("toy requested")

	s.toyStore.SetState(s.toyStore.toyRequested)

	return nil
}

func (s *HasToyState) GiveMoney(money int) error {
	return ErrSelectToyFirst
}

func (s *HasToyState) DispenseToy() error {
	return ErrSelectToyFirst
}
