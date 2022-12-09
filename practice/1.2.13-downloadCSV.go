package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/down", Welcome)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func GenerateCsv(filePath string, data [][]string) {
	fp, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	fp.WriteString("\xEF\xBB\xBF")
	writer := csv.NewWriter(fp)
	writer.WriteAll(data)
	writer.Flush()
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	filename := "exportUsers.csv"
	column := [][]string{
		{"手机号", "用户UID", "Email", "用户名"},
		{"18888888888", "2", "barry@163.com", "barry"},
		{"18888888888", "3", "wangwu@163.com", "wangwu"},
	}

	GenerateCsv(filename, column)

	file, err := os.Open(filename)
	content, err := ioutil.ReadAll(file)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("content-disposition", "attachment;filename=\""+filename+"\"")
	if err != nil {
		fmt.Println("Read File Err:", err.Error())
	} else {
		w.Write(content)
	}
}
