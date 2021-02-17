package circularprime

import (
	"fmt"

	"github.com/hsmtkk/project_euler/digitrotate"
	"github.com/hsmtkk/project_euler/prime"
)

func IsCircularPrime(n int64) bool {
	nums := digitrotate.DigitRotate(n)
	for _, n := range nums {
		if !prime.IsPrime(n) {
			return false
		}
	}
	return true
}

func Answer() int {
	var i int64
	ps := []int64{}
	for i = 2; i < 1000000; i++ {
		if IsCircularPrime(i) {
			ps = append(ps, i)
		}
	}
	fmt.Println(ps)
	return len(ps)
}
