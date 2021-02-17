package factdigitsum

import (
	"strconv"

	"github.com/hsmtkk/project_euler/factorial"
)

func Answer(n int) int {
	f := factorial.New().Factorial(n)
	ret := 0
	for _, d := range f.String() {
		n, _ := strconv.Atoi(string(d))
		ret += n
	}
	return ret
}
