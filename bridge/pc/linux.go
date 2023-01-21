package pc

import (
	"fmt"

	"github.com/IvanProg00/go-patterns/bridge/scanner"
)

type Linux struct {
	scanner scanner.Scanner
}

func (pc *Linux) AddScanner(sc scanner.Scanner) {
	pc.scanner = sc
}

func (pc *Linux) Scan() {
	fmt.Println("scan linux")
	pc.scanner.ScanFile()
}
