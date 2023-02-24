package chainofresponsibility

import "github.com/IvanProg00/go-patterns/chainofresponsibility/data"

func Run() {
	device := &data.Device{Name: "Device-1"}
	updateDataSvc := &data.UpdateService{Name: "Update-data-1"}
	dataSvc := &data.SaveService{}

	device.SetNext(updateDataSvc)
	updateDataSvc.SetNext(dataSvc)

	d := &data.Data{
		GetSource:    false,
		UpdateSource: false,
	}
	device.Execute(d)
}
