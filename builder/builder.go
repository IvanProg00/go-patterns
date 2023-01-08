package builder

import "github.com/IvanProg00/go-patterns/builder/computer"

func Run() {
	asusBuilder := computer.GetBuilder(computer.AsusType)
	lenovoBuilder := computer.GetBuilder(computer.LenovoType)

	factory := computer.NewFactory(asusBuilder)

	asus := factory.CreateComputer()
	asus.Print()

	factory.SetBuilder(lenovoBuilder)
	lenovo := factory.CreateComputer()
	lenovo.Print()
}
