package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Title string
	URL   string
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func main() {
	jsonStr := `{
		"data":{
			"children":[
					{
						"data":{
							"title":"Shidon,blog",
							"URL":"http://www.shidon.com"
						}
					}
				]
		}
	}`
	res := Response{}
	json.Unmarshal([]byte(jsonStr), &res)
	fmt.Println(res)
}
