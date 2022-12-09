package main

import (
	"fmt"
	"sort"
)

func main() {
	counts := map[string]int{"Barry": 96, "Aaron": 98, "Clan": 97}

	key := make([]string, 0, len(counts))
	for k := range counts {
		key = append(key, k)
	}

	sort.Slice(key, func(i, j int) bool { return counts[key[i]] > counts[key[j]] })

	for _, k := range key {
		fmt.Println(k, counts[k])
	}
}
