package digit

import (
	"math/big"
	"strconv"
)

func BigIntDigit(n *big.Int) []int {
	ds := []int{}
	for _, ch := range n.String() {
		n, _ := strconv.Atoi(string(ch))
		ds = append(ds, n)
	}
	return ds
}
