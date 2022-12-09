package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.Comma = ';'

	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}
