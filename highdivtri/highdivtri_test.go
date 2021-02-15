package highdivtri_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/highdivtri"
	"github.com/stretchr/testify/assert"
)

func TestAnswer1(t *testing.T) {
	want := 28
	got := highdivtri.Answer(5)
	assert.Equal(t, want, got)
}

func TestAnswer2(t *testing.T) {
	want := 76576500
	got := highdivtri.Answer(500)
	assert.Equal(t, want, got)
}
