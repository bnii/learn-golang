package arrays

func Sum(arr []int) (ret int) {
	ret = 0
	for _, num := range arr {
		ret += num
	}
	return
}

func SumAll(arrs ...[]int) (sums []int) {
	for _, numbers := range arrs {
		sums = append(sums, Sum(numbers))
	}
	return
}
