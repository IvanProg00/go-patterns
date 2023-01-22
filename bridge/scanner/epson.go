package scanner

import "fmt"

type Epson struct{}

func (s *Epson) ScanFile() {
	fmt.Println("Epson scans file")
}
