package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	plainStr := "Hello World."

	encodedStr := base64.StdEncoding.EncodeToString([]byte(plainStr))
	fmt.Println(encodedStr)

	decodedStrAsByteSlice, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		panic("Malformed input")
	}
	fmt.Println(string(decodedStrAsByteSlice))
}
