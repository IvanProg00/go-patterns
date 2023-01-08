package computer

type Lenovo struct {
	Brand       string
	Core        int
	Memory      int
	GraphicCard int
}

func (c *Lenovo) SetBrand() {
	c.Brand = "Lenovo"
}

func (c *Lenovo) SetCore() {
	c.Core = 8
}

func (c *Lenovo) SetMemory() {
	c.Memory = 16
}

func (c *Lenovo) SetGraphicCard() {
	c.GraphicCard = 2
}

func (c *Lenovo) GetComputer() Computer {
	return Computer{
		Brand:       c.Brand,
		Core:        c.Core,
		Memory:      c.Memory,
		GraphicCard: c.GraphicCard,
	}
}
