package route

type Route struct {
	Name       string
	TravelTime int
}

type Iterator interface {
	HasNext() bool
	GetNext() *Route
}

type Routes struct {
	Routes []Route
	index  int
}

func (r *Routes) HasNext() bool {
	return r.index < len(r.Routes)
}

func (r *Routes) GetNext() *Route {
	if r.HasNext() {
		val := &r.Routes[r.index]
		r.index++

		return val
	}

	return nil
}
