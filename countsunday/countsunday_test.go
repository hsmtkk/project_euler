package countsunday_test

import (
	"github.com/hsmtkk/project_euler/countsunday"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnswer(t *testing.T) {
	assert.Equal(t, 171, countsunday.Answer())
}
