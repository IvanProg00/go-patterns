package data

type Service interface {
	Execute(data *Data)
	SetNext(svc Service)
}

type Data struct {
	GetSource    bool
	UpdateSource bool
}
