package evenfib

import (
	"github.com/hsmtkk/project_euler/fibonacci"
	"github.com/hsmtkk/project_euler/sum"
)

func Answer() int {
	calc := fibonacci.New()
	ns := []int{}
	for n := 0; ; n++ {
		f := calc.Fibonacci(n)
		if f > 4000000 {
			break
		}
		if f%2 == 0 {
			ns = append(ns, f)
		}
	}
	return sum.Sum(ns)
}
