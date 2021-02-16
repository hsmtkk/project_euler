package factdigitsum_test

import (
	"github.com/hsmtkk/project_euler/factdigitsum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTen(t *testing.T) {
	assert.Equal(t, 27, factdigitsum.Answer(10))
}

func TestHundred(t *testing.T){
	assert.Equal(t, 648, factdigitsum.Answer(100))
}