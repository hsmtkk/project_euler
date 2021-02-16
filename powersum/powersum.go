package powersum

import(
	"math/big"
	"strconv"
)

func PowerSum(n int) int {
	ret := big.NewInt(1)
	for i:=0 ; i<n ; i++ {
		ret.Mul(ret, big.NewInt(2))
	}
	sum := 0
	for _, d := range ret.String() {
		n, _ := strconv.Atoi(string(d))
		sum += n
	}
	return sum
}