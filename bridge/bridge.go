package bridge

import (
	"github.com/IvanProg00/go-patterns/bridge/pc"
	"github.com/IvanProg00/go-patterns/bridge/scanner"
)

func Run() {
	// Scanners
	hp := &scanner.HP{}
	epson := &scanner.Epson{}

	// Computers
	var (
		linux   pc.PC = &pc.Linux{}
		windows pc.PC = &pc.Windows{}
		mac     pc.PC = &pc.Mac{}
	)

	linux.AddScanner(hp)
	linux.Scan()

	windows.AddScanner(epson)
	windows.Scan()

	mac.AddScanner(epson)
	mac.Scan()
}
