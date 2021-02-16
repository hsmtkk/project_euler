package perfectnum

import "github.com/hsmtkk/project_euler/sum"

func IsPerfectNumber(n int) bool {
	return n == sum.Sum(divisors(n))
}

func divisors(n int) []int {
	ds := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			ds = append(ds, i)
		}
	}
	return ds
}
