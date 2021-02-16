package lexperm_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/lexperm"
	"github.com/stretchr/testify/assert"
)

func TestLexPerm3(t *testing.T) {
	got := lexperm.LexPerm("012")
	want := []string{"012", "021", "102", "120", "201", "210"}
	assert.Equal(t, want, got)
}

func TestAnswer(t *testing.T) {
	want := "2783915460"
	got := lexperm.Answer()
	assert.Equal(t, want, got)
}
