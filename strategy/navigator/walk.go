package navigator

import "fmt"

type WalkStrategy struct{}

func (ps *WalkStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 4

	total := endPoint - startPoint
	totalTime := total * 60

	fmt.Printf("walk A %d to B %d; avg speed %d; total %d; total time %d min\n", startPoint, endPoint,
		avgSpeed, total, totalTime)
}
