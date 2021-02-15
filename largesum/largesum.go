package largesum

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func ReadDigits() ([]*big.Int, error) {
	fh, err := os.Open("digits.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open file; %w", err)
	}
	defer fh.Close()
	digits := []*big.Int{}
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		n := big.NewInt(0)
		n.SetString(line, 10)
		digits = append(digits, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan; %w", err)
	}
	return digits, nil
}

func Answer() (string, error) {
	ns, err := ReadDigits()
	if err != nil {
		return "", err
	}
	s := sum(ns).String()
	return s[0:10], nil
}

func sum(is []*big.Int) *big.Int {
	s := big.NewInt(0)
	for _, i := range is {
		s.Add(s, i)
	}
	return s
}
