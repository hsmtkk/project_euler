package mul35

func Mul35(num int) int {
	ns := Values35(num)
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}

func Values35(num int) []int {
	ns := []int{}
	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			ns = append(ns, i)
		}
	}
	return ns
}
