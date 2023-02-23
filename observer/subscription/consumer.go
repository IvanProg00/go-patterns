package subscription

type Consumer interface {
	Update(pubName string)
	GetName() string
}
