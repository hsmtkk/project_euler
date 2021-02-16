package namescore_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/namescore"
	"github.com/stretchr/testify/assert"
)

func TestNameValue(t *testing.T) {
	assert.Equal(t, 53, namescore.NameValue("COLIN"))
}

func TestReadNames(t *testing.T) {
	got, err := namescore.ReadNames()
	assert.Nil(t, err)
	assert.Equal(t, "COLIN", got[937])
}

func TestAnswer(t *testing.T) {
	want := 871198282
	got, err := namescore.Answer()
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
