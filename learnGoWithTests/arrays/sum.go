package arrays

func Sum(numbers []int) int {

	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		sums = append(sums, Sum(v))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, v := range numbersToSum {
		if len(v) != 0 {
			numbers := v[1:]
			sums = append(sums, Sum(numbers))
		} else {
			sums = append(sums, 0)
		}

	}
	return
}
