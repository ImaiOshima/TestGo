package main

import "fmt"

func main() {
	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		m[i] = 10 + i
	}

	for k := range m {
		if k == 5 {
			m[100] = 1000
		}
		delete(m, k)
		fmt.Println(k, m)
	}
}
