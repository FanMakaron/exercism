package differenceofsquares

func SquareOfSum(n int) int {
	res := 1
	for (n > 1){
		res += n
		n--
	}
	res *= res
	return res
}

func SumOfSquares(n int) int {
	res := 1
	for (n > 1){
		res += n*n
		n--
	}
	return res
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
