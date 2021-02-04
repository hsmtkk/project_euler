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
