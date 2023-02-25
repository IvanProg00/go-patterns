package navigator

import "fmt"

type RoadStrategy struct{}

func (rs *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30
	trafficJam := 2

	total := endPoint - startPoint
	totalTime := total * 40 * trafficJam

	fmt.Printf("road A %d to B %d; avg speed %d; traffic jam %d; total %d; total time %d min\n", startPoint, endPoint,
		avgSpeed, trafficJam, total, totalTime)
}
