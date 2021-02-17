package digitfactorial

import (
	"math/big"

	"github.com/hsmtkk/project_euler/digit"
	"github.com/hsmtkk/project_euler/factorial"
)

func Match(calc factorial.Calculator, n *big.Int) bool {
	sum := big.NewInt(0)
	for _, d := range digit.BigIntDigit(n) {
		sum.Add(sum, calc.Factorial(d))
	}
	return sum.Cmp(n) == 0
}

func Answer() int {
	calc := factorial.New()
	n := big.NewInt(10)
	upper := calc.Factorial(9)
	upper.Mul(upper, big.NewInt(7))
	sum := big.NewInt(0)
	for {
		if n.Cmp(upper) > 0 {
			break
		}
		if Match(calc, n) {
			sum.Add(sum, n)
		}
		n.Add(n, big.NewInt(1))
	}
	return int(sum.Int64())
}
