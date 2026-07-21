package darts

func square(x float64) float64 {
	return x * x
}

func Score(x, y float64) int {

	radiusTen := square(10)
	resultTen := float64(1)

	radiusFive := square(5)
	resultFive := float64(5)

	radiusOne := square(1)
	resultOne := float64(10)

	distance := square(x) + square(y)

	if distance <= radiusTen {
		if distance <= radiusFive {
			if distance <= radiusOne {
				return int(resultOne)
			}
			return int(resultFive)
		}
		return int(resultTen)
	}
	return 0
}
