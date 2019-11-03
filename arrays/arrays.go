package arrays

func Sum(arr []int) (ret int) {
	ret = 0
	for _, num := range arr {
		ret += num
	}
	return
}
