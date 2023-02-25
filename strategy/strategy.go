package strategy

import "github.com/IvanProg00/go-patterns/strategy/navigator"

const (
	start = 10
	end   = 100
)

func Run() {
	strategies := []navigator.Strategy{
		&navigator.PublicTransportStrategy{},
		&navigator.RoadStrategy{},
		&navigator.WalkStrategy{},
	}

	nav := navigator.Navigator{}

	for _, s := range strategies {
		nav.SetStrategy(s)
		nav.Route(start, end)
	}
}
