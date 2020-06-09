package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("hello")
	fmt.Println(sum(1, 2))
}

func sum(a int, b int) int {
	return a + b
}

func concatString(str ...string) string {
	ret := ""
	for _, ele := range str {
		ret += ele
	}
	return ret
}

func concatStringWithFor(str ...string) string {
	ret := ""
	length := len(str)
	for i := 0; i < length; i++ {
		ret += str[i]
	}
	return ret
}

func concatStringWithBuffer(str ...string) string {
	var buffer bytes.Buffer
	for _, ele := range str {
		buffer.WriteString(ele)
	}
	return buffer.String()
}

func concatStringWithNew(str ...string) string {
	buffer := new(bytes.Buffer)
	for _, ele := range str {
		buffer.WriteString(ele)
	}
	return buffer.String()
}
