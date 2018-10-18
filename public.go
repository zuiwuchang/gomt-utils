package utils

import (
	"github.com/zuiwuchang/gomt-utils/kk"
)

// Sum .
func Sum(vals ...int) (sum int) {
	for i := 0; i < len(vals); i++ {
		sum += vals[i]
	}
	kk.Out(sum)
	return
}
