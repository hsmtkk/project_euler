package mul35_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/mul35"
	"github.com/stretchr/testify/assert"
)

func TestMul35(t *testing.T) {
	want := 23
	got := mul35.Mul35(10)
	assert.Equal(t, want, got)
}

func TestValues35(t *testing.T) {
	want := []int{3, 5, 6, 9}
	got := mul35.Values35(10)
	assert.Equal(t, want, got)
}

func TestAnswer(t *testing.T) {
	want := 233168
	got := mul35.Mul35(1000)
	assert.Equal(t, want, got)
}
