package triangle_test

import (
	"testing"

	"github.com/hsmtkk/project_euler/triangle"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTriangleNumbers(t *testing.T) {
	want := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	got := []int{}
	ch := make(chan int)
	go triangle.GenerateTriangleNumbers(ch)
	for i := 0; i < len(want); i++ {
		n := <-ch
		got = append(got, n)
	}
	assert.Equal(t, want, got)
}
