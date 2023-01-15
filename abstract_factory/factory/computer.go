package factory

type IComputer interface {
	SetMemory(memory int)
	SetCPU(cpu int)
	GetMemory() int
	GetCPU() int
}

type Computer struct {
	Memory int
	CPU    int
}

func (c *Computer) SetMemory(memory int) {
	c.Memory = memory
}

func (c *Computer) SetCPU(cpu int) {
	c.CPU = cpu
}

func (c *Computer) GetMemory() int {
	return c.Memory
}

func (c *Computer) GetCPU() int {
	return c.CPU
}
