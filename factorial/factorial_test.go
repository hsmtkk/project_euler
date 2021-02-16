package factorial_test

import (
	"github.com/hsmtkk/project_euler/factorial"
	"github.com/stretchr/testify/assert"
	"testing"
	"math/big"
)

func TestTen(t *testing.T) {
	assert.Equal(t, big.NewInt(int64(3628800)), factorial.Factorial(10))
}
