package adapter

import (
	"github.com/IvanProg00/go-patterns/adapter/data"
	dataservice "github.com/IvanProg00/go-patterns/adapter/data_service"
)

func Run() {
	xmlDoc := data.XMLDocument{}
	dataservice.Send(&xmlDoc)

	jsonDoc := data.JSONDocument{}
	jsonAdapter := data.JSONDocumentAdapter{
		Doc: &jsonDoc,
	}
	dataservice.Send(&jsonAdapter)
}
