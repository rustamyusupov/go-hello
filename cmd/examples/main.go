package main

import (
	"fmt"
	example "go-hello/internal"
)

func main() {
	const rawResp = `
	{
			"header": {
					"code": 0,
					"message": ""
			},
			"data": [{
					"type": "user",
					"id": 100,
					"attributes": {
							"email": "bob@yandex.ru",
							"article_ids": [10, 11, 12]
					}
			}]
	}
	`

	resp, err := example.ReadResponse(rawResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)
}
