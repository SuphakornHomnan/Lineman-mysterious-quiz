package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, err := base64.StdEncoding.DecodeString(secret)

	if err != nil {
		fmt.Println("decode error: ", err)
	}

	decode := string(sd)
	reverse := func(str string) (result string) {
		for _, v := range str {
			result = string(v) + result
		}
		return
	}
	whatIsIt = reverse(decode)
	fmt.Println(whatIsIt)
}
