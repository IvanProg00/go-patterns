package pc

import "github.com/IvanProg00/go-patterns/bridge/scanner"

type PC interface {
	AddScanner(scanner scanner.Scanner)
	Scan()
}
