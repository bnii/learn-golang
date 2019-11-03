package arrays

func Sum(arr []int) (ret int) {
	ret = 0
	for _, num := range arr {
		ret += num
	}
	return
}

func SumAll(arrs ...[]int) (sums []int) {
	lengthOfNumbers := len(arrs)
	sums = make([]int, lengthOfNumbers)
	for i, numbers := range arrs {
		sums[i] = Sum(numbers)
	}
	return
}
