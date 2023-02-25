package navigator

import "fmt"

type PublicTransportStrategy struct{}

func (ps *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 40

	total := endPoint - startPoint
	totalTime := total * 40

	fmt.Printf("public transport A %d to B %d; avg speed %d; total %d; total time %d min\n", startPoint, endPoint,
		avgSpeed, total, totalTime)
}
