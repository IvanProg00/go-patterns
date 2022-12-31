package data

import "fmt"

type XMLDocument struct{}

func (doc XMLDocument) SendXMLData() {
	fmt.Println("<xml>xml document</xml>")
}
