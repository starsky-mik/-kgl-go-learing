package arrays_and_slices

func Sum(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func SumAll(arrays ...[]int) []int {
	var result []int

	for _, array := range arrays {
		result = append(result, Sum(array))
	}

	return result
}

func SumAllTails(arrays ...[]int) []int {
	var result []int

	for _, array := range arrays {
		sum := 0

		if len(array) > 1 {
			tail := array[1:]
			sum = Sum(tail)
		}

		result = append(result, sum)
	}

	return result
}
