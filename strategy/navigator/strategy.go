package navigator

type Strategy interface {
	Route(startPoint, endPoint int)
}
