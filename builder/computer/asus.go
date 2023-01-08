package computer

type Asus struct {
	Brand       string
	Core        int
	Memory      int
	GraphicCard int
}

func (c *Asus) SetBrand() {
	c.Brand = "Asus"
}

func (c *Asus) SetCore() {
	c.Core = 4
}

func (c *Asus) SetMemory() {
	c.Memory = 6
}

func (c *Asus) SetGraphicCard() {
	c.GraphicCard = 1
}

func (c *Asus) GetComputer() Computer {
	return Computer{
		Brand:       c.Brand,
		Core:        c.Core,
		Memory:      c.Memory,
		GraphicCard: c.GraphicCard,
	}
}
