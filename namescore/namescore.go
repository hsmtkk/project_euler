package namescore

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const namesTxt = "./names.txt"

func ReadNames() ([]string, error) {
	bs, err := ioutil.ReadFile(namesTxt)
	if err != nil {
		return nil, fmt.Errorf("failed to read file; %s; %w", namesTxt, err)
	}
	noQuote := strings.ReplaceAll(string(bs), "\"", "")
	splitted := strings.Split(noQuote, ",")
	sort.Strings(splitted)
	return splitted, nil
}

func NameValue(name string) int {
	v := 0
	for _, ch := range name {
		v += charValue(ch)
	}
	return v
}

func charValue(ch rune) int {
	alphas := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for index, alpha := range alphas {
		if ch == alpha {
			return index + 1
		}
	}
	return 0
}

func Answer() (int, error) {
	names, err := ReadNames()
	if err != nil {
		return 0, err
	}
	score := 0
	for i, name := range names {
		score += (i + 1) * NameValue(name)
	}
	return score, nil
}
