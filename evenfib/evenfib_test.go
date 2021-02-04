package evenfib_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/evenfib"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	want := 4613732
	got := evenfib.Answer()
	assert.Equal(t, want, got)
}
