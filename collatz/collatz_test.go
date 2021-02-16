package collatz_test

import (
	"github.com/hsmtkk/project_euler/collatz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollatz(t *testing.T) {
	assert.Equal(t, 40, collatz.Collatz(13))
	assert.Equal(t, 20, collatz.Collatz(40))
}

func TestSequence(t *testing.T) {
	want := []int{13, 40, 20, 10, 5, 16, 8, 4, 2, 1}
	got := collatz.Sequence(13)
	assert.Equal(t, want, got)
}
