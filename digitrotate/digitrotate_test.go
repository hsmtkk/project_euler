package digitrotate_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/digitrotate"
	"github.com/stretchr/testify/assert"
)

func TestDigitRotate(t *testing.T) {
	want := []int64{197, 971, 719}
	got := digitrotate.DigitRotate(197)
	assert.Equal(t, want, got)
}
