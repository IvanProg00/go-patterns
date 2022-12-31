package dataservice

type AnalyticalDataService interface {
	SendXMLData()
}

func Send(svc AnalyticalDataService) {
	svc.SendXMLData()
}
