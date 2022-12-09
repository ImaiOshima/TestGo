package main

import "fmt"

func main() {
	m := map[string]string{
		"a": "value_a",
		"b": "value_b",
	}

	var sli []*string
	for k, v := range m {
		fmt.Println("k:[%p],v:[%p]", &k, &v)
		sli = append(sli, &v)
	}
	for _, b := range sli {
		fmt.Println(*b)
	}
}
