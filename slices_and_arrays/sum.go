package slicesandarrays

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sums = append(sums, Sum(nums))
		} else {
			tails := nums[1:]
			sums = append(sums, Sum(tails))
		}
	}

	return sums
}
