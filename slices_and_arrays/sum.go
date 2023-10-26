package slicesandarrays

func Sum(numbers [5]int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}
