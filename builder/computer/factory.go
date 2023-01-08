package computer

type Factory struct {
	b Builder
}

func NewFactory(b Builder) *Factory {
	return &Factory{
		b: b,
	}
}

func (f *Factory) SetBuilder(b Builder) {
	f.b = b
}

func (f *Factory) CreateComputer() Computer {
	f.b.SetBrand()
	f.b.SetCore()
	f.b.SetGraphicCard()
	f.b.SetMemory()

	return f.b.GetComputer()
}
