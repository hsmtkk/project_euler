package factorial_test

import (
	"math/big"
	"testing"

	"github.com/hsmtkk/project_euler/factorial"
	"github.com/stretchr/testify/assert"
)

func TestTen(t *testing.T) {
	calculator := factorial.New()
	assert.Equal(t, big.NewInt(int64(3628800)), calculator.Factorial(10))
}
