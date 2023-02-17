package toystore

type State interface {
	AddToys(count int) error
	RequestToy() error
	GiveMoney(money int) error
	DispenseToy() error
}
