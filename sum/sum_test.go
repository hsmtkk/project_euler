package sum_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/sum"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	ns := []int{10, 20, 30}
	want := 60
	got := sum.Sum(ns)
	assert.Equal(t, want, got)
}
