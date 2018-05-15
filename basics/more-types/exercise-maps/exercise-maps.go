package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, key := range strings.Fields(s) {
		if _, ok := m[key]; !ok {
			m[key] = 1
		} else {
			m[key]++
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
