# Adapter

The Adapter pattern allows the interface of an existing class to be used as
another interface.

## Example

We will create an analytical service, that will send documents in XML format.

1. Create Interface and function that will send document.

```go
// data_service/analytical.go

type AnalyticalDataService interface {
	SendXMLData()
}

func Send(svc AnalyticalDataService) {
	svc.SendXMLData()
}
```

2. Now we will create an **XML** document.

```go
// data/xml.go

type XMLDocument struct{}

func (doc XMLDocument) SendXMLData() {
	fmt.Println("<xml>xml document</xml>")
}
```

3. It works, but now we want to send **JSON** documents. So let's create **JSON** document.

```go
// data/json.go

type JSONDocument struct{}

func (doc JSONDocument) ConvertToXML() string {
	return "<xml>json document</xml>"
}
```

4. But our analytical service can not send **JSON** files, because it only knows how to work with **XML**. So we will create adapter for it using `ConvertToXML`.

```go
// data/json.go

type JSONDocumentAdapter struct {
	Doc *JSONDocument
}

func (adapter JSONDocumentAdapter) SendXMLData() {
	fmt.Println(adapter.Doc.ConvertToXML())
}
```

5. Let's execute our app.

```bash
<xml>xml document</xml>
<xml>json document</xml>
```
