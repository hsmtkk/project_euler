package circularprime_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/circularprime"
	"github.com/stretchr/testify/assert"
)

func TestIsCircularPrime(t *testing.T) {
	assert.False(t, circularprime.IsCircularPrime(196))
	assert.True(t, circularprime.IsCircularPrime(197))
	assert.False(t, circularprime.IsCircularPrime(198))
}

func TestAnswer(t *testing.T) {
	want := 55
	got := circularprime.Answer()
	assert.Equal(t, want, got)
}
