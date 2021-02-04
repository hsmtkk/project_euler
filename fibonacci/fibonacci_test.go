package fibonacci_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/fibonacci"
	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	want := 89
	got := fibonacci.New().Fibonacci(9)
	assert.Equal(t, want, got)
}
