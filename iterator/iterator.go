package iterator

import (
	"fmt"

	"github.com/IvanProg00/go-patterns/iterator/route"
)

func Run() {
	routes := route.Routes{
		Routes: []route.Route{
			{Name: "Route-1", TravelTime: 110},
			{Name: "Route-2", TravelTime: 50},
			{Name: "Route-3", TravelTime: 60},
			{Name: "Route-4", TravelTime: 30},
		},
	}

	count := 1

	for routes.HasNext() {
		r := routes.GetNext()
		fmt.Printf("%d. %s -> time %d\n", count, r.Name, r.TravelTime)

		count++
	}
}
