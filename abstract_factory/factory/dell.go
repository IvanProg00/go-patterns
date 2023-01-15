package factory

type Dell struct{}

func (d *Dell) GetComputer() IComputer {
	return &DellComputer{
		Computer: Computer{
			Memory: 8,
			CPU:    4,
		},
	}
}

func (d *Dell) GetMonitor() IMonitor {
	return &DellMonitor{
		Monitor: Monitor{
			Size: 27,
		},
	}
}
