package pc

type HomePC struct {
	Wrapper      Wrapper
	CPU          int
	GraphicsCard int
}

func (pc HomePC) GetPrice() float64 {
	return pc.Wrapper.GetPrice() * float64(pc.CPU) * float64(pc.GraphicsCard)
}
