package largestprime_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/largestprime"
	"github.com/stretchr/testify/assert"
)

func TestLargestPrime(t *testing.T) {
	want := 29
	got := largestprime.LargestPrime(13195)
	assert.Equal(t, want, got)
}

func TestAnswer(t *testing.T) {
	want := 6857
	got := largestprime.LargestPrime(600851475143)
	assert.Equal(t, want, got)
}
