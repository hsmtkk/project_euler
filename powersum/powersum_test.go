package powersum_test

import (
	"github.com/hsmtkk/project_euler/powersum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T){
	assert.Equal(t, 26, powersum.PowerSum(15))
}

func TestAnswer(t *testing.T){
	assert.Equal(t, 1366, powersum.PowerSum(1000))
}
