package vehicle

type Vehicle interface {
	Arrive()
	Go()
	PermitArrive()
}

type Dispatcher interface {
	CanArrive(v Vehicle) bool
	NotifyAboutGo()
}
