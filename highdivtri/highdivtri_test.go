package highdivtri_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/highdivtri"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	want := 76576500
	got := highdivtri.Answer()
	assert.Equal(t, want, got)
}
