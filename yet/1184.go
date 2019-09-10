var D []int

func calcCW(src, dst int) int {
	res := 0
	for i := src; i != dst; i++ {
		if i == len(D) {
			i = 0
		}
		res += D[i]
	}
	return res
}

func calcCCW(src, dst int) int {
	res := 0
	for i := src - 1; i + 1 != dst; i-- {
		if i < 0 {
			i = len(D) - 1
		}
		res += D[i]
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start == destination {
		return 0
	}

	D = distance

	// clock wise
	cw := calcCW(start, destination)
	// counter clock wise
	ccw := calcCCW(start, destination)

	return min(cw, ccw)
}