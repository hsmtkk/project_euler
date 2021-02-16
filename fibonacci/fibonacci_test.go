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

func TestGenerator(t *testing.T) {
	ch := make(chan int64)
	go fibonacci.Generate(ch)
	want := []int64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	got := []int64{}
	for i := 0; i < 10; i++ {
		n := <-ch
		got = append(got, n)
	}
	assert.Equal(t, want, got)
}
