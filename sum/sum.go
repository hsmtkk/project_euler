package sum

func Sum(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}

func Sum64(ns []int64) int64 {
	s := int64(0)
	for _, n := range ns {
		s += n
	}
	return s
}
