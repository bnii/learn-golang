package arrays

func Sum(arr []int) (ret int) {
	ret = 0
	for _, num := range arr {
		ret += num
	}
	return
}

func SumAll(arrs ...[]int) (sums []int) {
	return SumOmitFirstX(0, arrs)
}

func SumAllTails(ints ...[]int) (sums []int) {
	return SumOmitFirstX(1, ints)
}

func SumOmitFirstX(x int, ints [][]int) (sums []int) {
	for _, numbers := range ints {
		var nextAdd int
		if len(numbers) < x {
			nextAdd = 0;
		} else {
			nextAdd = Sum(numbers[x:])
		}
		sums = append(sums, nextAdd)
	}
	return
}
