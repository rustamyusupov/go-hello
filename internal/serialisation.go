package example

import "encoding/json"

func Serialise(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
