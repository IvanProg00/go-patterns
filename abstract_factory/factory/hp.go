package factory

type HP struct{}

func (h *HP) GetMonitor() IMonitor {
	return &HPMonitor{
		Monitor: Monitor{
			Size: 24,
		},
	}
}

func (h *HP) GetComputer() IComputer {
	return &HPComputer{
		Computer: Computer{
			Memory: 16,
			CPU:    8,
		},
	}
}
