package highdivtri

import (
	"github.com/hsmtkk/project_euler/divisors"
	"github.com/hsmtkk/project_euler/triangle"
)

func Answer(numOfDiv int) int {
	ch := make(chan int)
	go triangle.GenerateTriangleNumbers(ch)
	for t := range ch {
		ds := divisors.Divisors(t)
		if len(ds) > numOfDiv {
			return t
		}
	}
	return 0
}
