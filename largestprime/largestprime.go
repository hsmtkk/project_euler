package largestprime

import (
	"github.com/hsmtkk/project_euler/prime"
)

func LargestPrime(n int) int {
	ps := prime.Primes(n)
	return ps[len(ps)-1]
}
