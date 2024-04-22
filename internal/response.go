package example

import (
	"encoding/json"
	"fmt"
)

type Header struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Type       string `json:"type"`
	ID         int    `json:"id"`
	Attributes struct {
		Email      string `json:"email"`
		ArticleIDs []int  `json:"article_ids"`
	} `json:"attributes"`
}

type Response struct {
	Header Header `json:"header"`
	Data   []Data `json:"data"`
}

func ReadResponse(rawResp string) (Response, error) {
	resp := Response{}

	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
	}

	return resp, nil
}
