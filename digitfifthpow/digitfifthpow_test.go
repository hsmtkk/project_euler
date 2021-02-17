package digitfifthpow_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/digitfifthpow"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	assert.Equal(t, int64(443839), digitfifthpow.Answer())
}
