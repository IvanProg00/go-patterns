package pc

type BasePC struct {
	CPU int
}

func (pc BasePC) GetPrice() float64 {
	return 10.0
}
