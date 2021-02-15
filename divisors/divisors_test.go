package divisors_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/divisors"
	"github.com/stretchr/testify/assert"
)

func TestDivisors(t *testing.T) {
	want := []int{1, 2, 4, 7, 14, 28}
	got := divisors.Divisors(28)
	assert.Equal(t, want, got)
}
