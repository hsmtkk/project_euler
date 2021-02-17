package digit_test

import (
	"math/big"
	"testing"

	"github.com/hsmtkk/project_euler/digit"
	"github.com/stretchr/testify/assert"
)

func TestBigIntDigit(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := digit.BigIntDigit(big.NewInt(123456789))
	assert.Equal(t, want, got)
}
