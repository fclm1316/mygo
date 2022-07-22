package tool

import (
	"encoding/json"
	// "fmt"
	"io"
)

type JsonParse struct {
}

func Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}

func SliceToByte(v interface{}) ([]byte, error) {
	jsonbyte, error := json.Marshal(v)
	return jsonbyte, error

}

func ByteToJosn(b []byte) []map[string]interface{} {
	var data []map[string]interface{}
	json.Unmarshal(b, &data)
	return data

}
