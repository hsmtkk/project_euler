package longestcollatz

import(
	"github.com/hsmtkk/project_euler/collatz"
)

func Answer() int {
	maxNum := 0
	maxLength := 0
	for i := 1; i < 1000000; i++ {
		seq := collatz.Sequence(i)
		if len(seq) > maxLength {
			maxNum = i
			maxLength = len(seq)
		}
	}
	return maxNum
}
