package longestcollatz_test

import(
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/hsmtkk/project_euler/longestcollatz"
)

func TestAnswer(t *testing.T) {
	want := 837799
	got := longestcollatz.Answer()
	assert.Equal(t, want, got)
}
