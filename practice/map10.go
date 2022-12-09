package main

import (
	"fmt"
	"sort"
)

type frequency struct {
	char string
	fre  int
}

func main() {
	str := "hiilovegogogo"
	fmt.Println(str)
	f := frequencies(str)
	fmt.Println(f)
}

func frequencies(s string) []frequency {
	m := make(map[string]int)
	for _, r := range s {
		m[string(r)]++
	}
	a := make([]frequency, 0, len(m))
	for c, f := range m {
		a = append(a, frequency{char: c, fre: f})
	}
	sort.Slice(a, func(i, j int) bool { return a[i].fre > a[j].fre })
	return a
}
