package subscription

type Subject interface {
	Subscribe(c Consumer)
	Unsubscribe(c Consumer)
	Notify()
}
