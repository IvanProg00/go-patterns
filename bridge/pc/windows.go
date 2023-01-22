package pc

import (
	"fmt"

	"github.com/IvanProg00/go-patterns/bridge/scanner"
)

type Windows struct {
	scanner scanner.Scanner
}

func (pc *Windows) AddScanner(sc scanner.Scanner) {
	pc.scanner = sc
}

func (pc *Windows) Scan() {
	fmt.Println("scan windows")
	pc.scanner.ScanFile()
}
