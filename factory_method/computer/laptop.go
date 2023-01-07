package computer

type Laptop struct {
	Type Type
	OS   string
	CPU  int
}

func NewLaptop() Computer {
	return &Laptop{
		Type: LaptopType,
		OS:   "MacOS",
		CPU:  8,
	}
}

func (l *Laptop) GetType() string {
	return string(l.Type)
}

func (l *Laptop) GetCPU() int {
	return l.CPU
}

func (l *Laptop) GetOS() string {
	return l.OS
}
