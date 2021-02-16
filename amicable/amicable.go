package amicable

import "github.com/hsmtkk/project_euler/sum"

func Amicable(n int) (int, bool) {
	s := sum.Sum(divisors(n))
	if n == s {
		return 0, false
	}
	m := sum.Sum(divisors(s))
	if m == n {
		return s, true
	}
	return 0, false
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

func Answer() int {
	as := []int{}
	for i := 1; i < 10000; i++ {
		if _, found := Amicable(i); found {
			as = append(as, i)
		}
	}
	return sum.Sum(as)
}
