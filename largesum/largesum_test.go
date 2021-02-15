package largesum_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/largesum"
	"github.com/stretchr/testify/assert"
)

func TestReadDigits(t *testing.T) {
	digits, err := largesum.ReadDigits()
	assert.Nil(t, err)
	assert.Equal(t, 100, len(digits))
}

func TestAnswer(t *testing.T) {
	want := "5537376230"
	got, err := largesum.Answer()
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
