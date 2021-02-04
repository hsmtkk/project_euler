package util

func Sum(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}
