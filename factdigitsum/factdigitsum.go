package factdigitsum

import (
	"github.com/hsmtkk/project_euler/factorial"
	"strconv"
)

func Answer(n int) int {
	f := factorial.Factorial(n)
	ret := 0
	for _, d := range f.String() {
		n, _ := strconv.Atoi(string(d))
		ret += n
	}
	return ret
}
