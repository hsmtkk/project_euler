package fibonacci

type FibonacciCalculator interface {
	Fibonacci(n int) int
}

type calculatorImpl struct {
	memo map[int]int
}

func New() FibonacciCalculator {
	return &calculatorImpl{memo: map[int]int{}}
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
