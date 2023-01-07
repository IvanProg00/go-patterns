package computer

import "errors"

type Type string

type Computer interface {
	GetType() string
	GetCPU() int
	GetOS() string
}

const (
	LaptopType           Type = "laptop"
	PersonalComputerType Type = "personal computer"
	ServerType           Type = "server"
)

var ErrComputerNotFound = errors.New("computer not found")

func New(compType Type) (Computer, error) {
	switch compType {
	case LaptopType:
		return NewLaptop(), nil
	case PersonalComputerType:
		return NewPersonalComputer(), nil
	case ServerType:
		return NewServer(), nil
	default:
		return nil, ErrComputerNotFound
	}
}
