package computer

type Builder interface {
	SetBrand()
	SetCore()
	SetMemory()
	SetGraphicCard()
	GetComputer() Computer
}

type Type int

const (
	LenovoType Type = iota
	AsusType
)

func GetBuilder(builderType Type) Builder {
	switch builderType {
	case LenovoType:
		return &Lenovo{}
	case AsusType:
		return &Asus{}
	default:
		return nil
	}
}
