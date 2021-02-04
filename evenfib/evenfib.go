package evenfib

import "github.com/hsmtkk/project_euler/util"

type FibonacciCalculator interface {
	Fibonacci(n int) int
}

type calculatorImpl struct {
	memo map[int]int
}

func New() FibonacciCalculator {
	memo := map[int]int{}
	return &calculatorImpl{memo: memo}
}

func (c *calculatorImpl) Fibonacci(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	ans, ok := c.memo[n]
	if ok {
		return ans
	}
	ans = c.Fibonacci(n-1) + c.Fibonacci(n-2)
	c.memo[n] = ans
	return ans
}

func Answer() int {
	calc := New()
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
	return util.Sum(ns)
}
