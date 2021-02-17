package digitfactorial_test

import (
	"math/big"
	"testing"

	"github.com/hsmtkk/project_euler/digitfactorial"
	"github.com/hsmtkk/project_euler/factorial"
	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	calc := factorial.New()
	assert.False(t, digitfactorial.Match(calc, big.NewInt(144)))
	assert.True(t, digitfactorial.Match(calc, big.NewInt(145)))
	assert.False(t, digitfactorial.Match(calc, big.NewInt(146)))
}

func TestAnswer(t *testing.T) {
	assert.Equal(t, 40730, digitfactorial.Answer())
}
