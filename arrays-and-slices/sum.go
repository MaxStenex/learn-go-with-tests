package arrays_and_slices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	sums := make([]int, len(slices))

	for i, slice := range slices {
		sums[i] = Sum(slice)
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	var sumsOfTails []int

	for _, slice := range slices {
		if len(slice) == 0 {
			continue
		}

		tailElement := slice[len(slice)-1]
		sumsOfTails = append(sumsOfTails, tailElement)
	}

	return sumsOfTails
}
