package triangle

func GenerateTriangleNumbers(ch chan<- int) {
	s := 0
	for i := 1; ; i++ {
		s += i
		ch <- s
	}
}
