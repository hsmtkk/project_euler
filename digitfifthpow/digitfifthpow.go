package digitfifthpow

import (
	"math"
	"strconv"

	"github.com/hsmtkk/project_euler/sum"
)

func Answer() int64 {
	ns := []int64{}
	var i int64
	for i = 2; i < 1000000; i++ {
		if match(i) {
			ns = append(ns, i)
		}
	}
	return sum.Sum64(ns)
}

func pow5(n int64) int64 {
	return int64(math.Pow(float64(n), 5))
}

func match(n int64) bool {
	s := int64(0)
	for _, d := range strconv.FormatInt(n, 10) {
		m, _ := strconv.ParseInt(string(d), 10, 64)
		s += pow5(m)
	}
	return n == s
}
