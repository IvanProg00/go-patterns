package factory

import "errors"

type ComputerFactory interface {
	GetMonitor() IMonitor
	GetComputer() IComputer
}

type Brand string

const (
	DellBrand Brand = "dell"
	HPBrand   Brand = "hp"
)

var ErrBrandNotFound = errors.New("brand not found")

func GetComputerFactory(brand Brand) (ComputerFactory, error) {
	switch brand {
	case DellBrand:
		return &Dell{}, nil
	case HPBrand:
		return &HP{}, nil
	default:
		return nil, ErrBrandNotFound
	}
}
