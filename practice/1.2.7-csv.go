package main

import (
	"encoding/csv"
	"os"
)

func main() {
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	w.Write([]string{"学号", "姓名", "分数"})
	w.Write([]string{"1", "Barry", "99.5"})
	w.Write([]string{"2", "Sharon", "100"})
	w.Flush()
}
