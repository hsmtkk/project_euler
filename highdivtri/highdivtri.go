package highdivtri

import (
	"fmt"

	"github.com/hsmtkk/project_euler/divisors"
	"github.com/hsmtkk/project_euler/triangle"
)

func Answer() int {
	ch := make(chan int)
	go triangle.GenerateTriangleNumbers(ch)
	for t := range ch {
		ds := divisors.Divisors(t)
		if len(ds)%100 == 0 {
			fmt.Println(t)
		}
		if len(ds) > 500 {
			return t
		}
	}
	return 0
}
