package data

import "fmt"

type JSONDocument struct{}

type JSONDocumentAdapter struct {
	Doc *JSONDocument
}

func (doc JSONDocument) ConvertToXML() string {
	return "<xml>json document</xml>"
}

func (adapter JSONDocumentAdapter) SendXMLData() {
	fmt.Println(adapter.Doc.ConvertToXML())
}
