package thousandfibonacci_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/thousandfibonacci"
	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	assert.Equal(t, 4782, thousandfibonacci.Answer())
}
