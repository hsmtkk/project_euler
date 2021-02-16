package perfectnum_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/perfectnum"
	"github.com/stretchr/testify/assert"
)

func TestPerfectNum(t *testing.T) {
	assert.True(t, perfectnum.IsPerfectNumber(28))
	assert.False(t, perfectnum.IsPerfectNumber(27))
}
