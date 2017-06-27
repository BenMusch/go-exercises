package diffsquares

func SumOfSquares(upperBound int) int {
	sum := 0

	for _, num := range NumRange(upperBound) {
		sum += num * num
	}
	return sum
}

func SquareOfSums(upperBound int) int {
	sum := 0

	for _, num := range NumRange(upperBound) {
		sum += num
	}
	return sum * sum
}

func Difference(upperBound int) int {
	return SquareOfSums(upperBound) - SumOfSquares(upperBound)
}

func NumRange(upperBound int) []int {
	ints := make([]int, upperBound)

	for i := 1; i <= upperBound; i++ {
		ints[i-1] = i
	}
	return ints
}
