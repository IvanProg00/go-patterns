package toystore

type NoToysState struct {
	toyStore *ToyStore
}

func (s *NoToysState) AddToys(count int) error {
	s.toyStore.IncrementItemCount(count)
	s.toyStore.SetState(s.toyStore.hasToy)

	return nil
}

func (s *NoToysState) RequestToy() error {
	return ErrNoToys
}

func (s *NoToysState) GiveMoney(money int) error {
	return ErrNoToys
}

func (s *NoToysState) DispenseToy() error {
	return ErrNoToys
}
