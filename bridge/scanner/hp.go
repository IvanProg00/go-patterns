package scanner

import "fmt"

type HP struct{}

func (s *HP) ScanFile() {
	fmt.Println("HP scans file")
}
