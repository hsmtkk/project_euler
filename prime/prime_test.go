package prime_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/prime"
	"github.com/stretchr/testify/assert"
)

func TestZero(t *testing.T) {
	want := []int{5, 7, 13, 29}
	got := prime.Primes(13195)
	assert.Equal(t, want, got)
}

func TestOne(t *testing.T) {
	want := []int{2, 2, 2, 2}
	got := prime.Primes(16)
	assert.Equal(t, want, got)
}

func TestIsPrime(t *testing.T) {
	assert.False(t, prime.IsPrime(996))
	assert.True(t, prime.IsPrime(997))
	assert.False(t, prime.IsPrime(998))
	assert.False(t, prime.IsPrime(999))
}
