package factorial

import (
	"math/big"
)

type Calculator interface {
	Factorial(n int) *big.Int
}

type calculatorImpl struct {
	memo map[int]*big.Int
}

func New() Calculator {
	return &calculatorImpl{memo: map[int]*big.Int{}}
}

func (c *calculatorImpl) Factorial(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(1))
	}
	ret, ok := c.memo[n]
	if ok {
		return ret
	}
	ret = big.NewInt(int64(n))
	ret.Mul(ret, c.Factorial(n-1))
	c.memo[n] = ret
	return ret
}
