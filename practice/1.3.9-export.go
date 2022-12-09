package main

import (
	"database/sql"
	"fmt"
)

var (
	tables = []string{"user", "order"}
	count  = len(tables)
	ch     = make(chan bool, count)
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/chapter1?charset=utf-8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for _, table := range tables {
		go SqlQuery(db, table, ch)
	}

	for i := 0; i < count; i++ {
		<-ch
	}

	fmt.Println("完成！")
}

func SqlQuery(db *sql.DB, table string, ch chan bool) {

}
