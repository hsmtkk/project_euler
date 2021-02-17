package digitrotate

import "strconv"

func DigitRotate(n int64) []int64 {
	results := []int64{}
	ds := strconv.FormatInt(n, 10)
	for i := 0; i < len(ds); i++ {
		left := ds[:i]
		right := ds[i:]
		s := right + left
		m, _ := strconv.ParseInt(s, 10, 64)
		results = append(results, m)
	}
	return results
}
