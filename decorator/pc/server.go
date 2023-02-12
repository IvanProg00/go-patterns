package pc

type ServerPC struct {
	Wrapper Wrapper
	CPU     int
	Memory  int
}

func (pc ServerPC) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.CPU) * float64(pc.Memory)
}
