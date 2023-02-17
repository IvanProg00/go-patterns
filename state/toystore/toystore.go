package toystore

import "fmt"

type ToyStore struct {
	hasToy       State
	toyRequested State
	hasMoney     State
	noToys       State
	currentState State

	toysCount int
	toysPrice int
}

func New(toysCount, toysPrice int) *ToyStore {
	toyStore := &ToyStore{
		toysCount: toysCount,
		toysPrice: toysPrice,
	}

	toyStore.hasToy = &HasToyState{toyStore: toyStore}
	toyStore.toyRequested = &ToyRequestedState{toyStore: toyStore}
	toyStore.hasMoney = &HasMoneyState{ToyStore: toyStore}
	toyStore.noToys = &NoToysState{toyStore: toyStore}

	toyStore.SetState(toyStore.hasToy)

	return toyStore
}

func (t *ToyStore) AddToys(count int) error {
	return t.currentState.AddToys(count)
}

func (t *ToyStore) RequestToy() error {
	return t.currentState.RequestToy()
}

func (t *ToyStore) GiveMoney(money int) error {
	return t.currentState.GiveMoney(money)
}

func (t *ToyStore) DispenseToy() error {
	return t.currentState.DispenseToy()
}

func (t *ToyStore) SetState(s State) {
	t.currentState = s
}

func (t *ToyStore) IncrementItemCount(count int) {
	fmt.Printf("add toys %d\n", count)

	t.toysCount += count
}
