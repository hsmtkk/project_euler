package thousandfibonacci

import (
	"strconv"

	"github.com/hsmtkk/project_euler/fibonacci"
)

func Answer() int {
	ch := make(chan int64)
	go fibonacci.Generate(ch)
	for i := 1; ; i++ {
		f := <-ch
		d := strconv.FormatInt(f, 10)
		if len(d) >= 1000 {
			return i
		}
	}
}
