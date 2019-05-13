package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := []byte("123")
	m := map[string]interface{}{
		"key": string(a),
	}

	mByte, _ := json.Marshal(m)
	fmt.Println(mByte)

	fmt.Println(string(mByte))
}