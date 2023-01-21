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
	linux := &pc.Linux{}
	windows := &pc.Windows{}
	mac := &pc.Mac{}

	linux.AddScanner(hp)
	linux.Scan()

	windows.AddScanner(epson)
	windows.Scan()

	mac.AddScanner(epson)
	mac.Scan()
}
