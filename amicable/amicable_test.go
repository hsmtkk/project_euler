package amicable_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/amicable"
	"github.com/stretchr/testify/assert"
)

func TestAmicable(t *testing.T) {
	want := 284
	got, found := amicable.Amicable(220)
	assert.True(t, found)
	assert.Equal(t, want, got)

	want = 220
	got, found = amicable.Amicable(284)
	assert.True(t, found)
	assert.Equal(t, want, got)

	_, found = amicable.Amicable(221)
	assert.False(t, found)
}

func TestAnswer(t *testing.T) {
	assert.Equal(t, 31626, amicable.Answer())
}
