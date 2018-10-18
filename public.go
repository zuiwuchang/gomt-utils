package utils

// Sum .
func Sum(vals ...int) (sum int) {
	for i := 0; i < len(vals); i++ {
		sum += vals[i]
	}
	return
}
