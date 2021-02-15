package divisors

func Divisors(n int) []int {
	ds := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ds = append(ds, i)
		}
	}
	return ds
}
