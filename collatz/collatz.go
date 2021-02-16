package collatz

func Collatz(n int) int {
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return n / 2
	}
	if n%2 == 1 {
		return 3*n + 1
	}
	return 1
}

func Sequence(n int) []int {
	ret := []int{}
	for {
		ret = append(ret, n)
		if n == 1 {
			break
		}
		n = Collatz(n)
	}
	return ret
}
