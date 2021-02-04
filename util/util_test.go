package util_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/util"
	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	ns := []int{10, 20, 30}
	want := 60
	got := util.Sum(ns)
	assert.Equal(t, want, got)
}
