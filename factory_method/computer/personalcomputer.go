package computer

type PersonalComputer struct {
	Type Type
	OS   string
	CPU  int
}

func NewPersonalComputer() Computer {
	return &PersonalComputer{
		Type: PersonalComputerType,
		OS:   "Windows",
		CPU:  16,
	}
}

func (p *PersonalComputer) GetType() string {
	return string(p.Type)
}

func (p *PersonalComputer) GetCPU() int {
	return p.CPU
}

func (p *PersonalComputer) GetOS() string {
	return p.OS
}
