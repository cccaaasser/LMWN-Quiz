package main

import (
	"encoding/base64"
	"fmt"
)

// function to reverse string
func reverseString(str string) (result string) {
	// iterate over str and prepend to result
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)

	fmt.Printf("Answer= %s\n", sd)

	//readable
	whatIsIt = reverseString(string(sd))
	fmt.Printf("Readable= %s\n", whatIsIt)
}
