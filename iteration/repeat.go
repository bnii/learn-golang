package iteration

func Repeat(c string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += c
	}
	return repeated

}
