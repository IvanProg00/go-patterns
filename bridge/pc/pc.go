package pc

import "text/scanner"

type PC interface {
	AddScanner(scanner scanner.Scanner)
	Scan()
}
