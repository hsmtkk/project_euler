package prime

import "math"

func Primes(n int) []int {
	m := n
	ps := []int{}
	sq := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq; i++ {
		for {
			if m%i != 0 {
				break
			}
			m /= i
			ps = append(ps, i)
		}
	}
	return ps
}

func IsPrime(n int64) bool {
	var i int64
	sq := int64(math.Sqrt(float64(n)))
	for i = 2; i <= sq; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
