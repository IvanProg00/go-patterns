package pc

import (
	"fmt"

	"github.com/IvanProg00/go-patterns/bridge/scanner"
)

type Mac struct {
	scanner scanner.Scanner
}

func (pc *Mac) AddScanner(sc scanner.Scanner) {
	pc.scanner = sc
}

func (pc *Mac) Scan() {
	fmt.Println("scan mac")
	pc.scanner.ScanFile()
}
