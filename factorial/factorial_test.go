package factorial_test

import (
	"math/big"
	"testing"

	"github.com/hsmtkk/project_euler/factorial"
	"github.com/stretchr/testify/assert"
)

func TestTen(t *testing.T) {
	assert.Equal(t, big.NewInt(int64(3628800)), factorial.Factorial(10))
}
