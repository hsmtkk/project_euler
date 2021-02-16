package lexperm

import "sort"

func LexPerm(digits string) []string {
	if len(digits) == 1 {
		return []string{digits}
	}
	if len(digits) == 2 {
		return []string{digits, string(digits[1]) + string(digits[0])}
	}
	results := []string{}
	for i := range digits {
		left := digits[:i]
		right := digits[i+1:]
		for _, rec := range LexPerm(left + right) {
			results = append(results, string(digits[i])+rec)
			results = append(results, rec+string(digits[i]))
		}
	}
	return uniqued(results)
}

func uniqued(orig []string) []string {
	m := map[string]bool{}
	for _, s := range orig {
		m[s] = true
	}
	results := []string{}
	for s := range m {
		results = append(results, s)
	}
	sort.Strings(results)
	return results
}

func Answer() string {
	ds := LexPerm("0123456789")
	return ds[1000000-1]
}
