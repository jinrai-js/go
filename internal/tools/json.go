package tools

import (
	"encoding/json"
	"io"
)

func StrToJson(from string) any {
	var result any
	json.Unmarshal([]byte(from), &result)

	return result
}

func IoToJson(stream io.ReadCloser) any {
	var result any
	json.NewDecoder(stream).Decode(&result)

	return result
}
