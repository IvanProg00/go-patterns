package factory

type IMonitor interface {
	SetSize(size int)
	GetSize() int
}

type Monitor struct {
	Size int
}

func (m *Monitor) SetSize(size int) {
	m.Size = size
}

func (m *Monitor) GetSize() int {
	return m.Size
}
