package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))
	z := fmt.Sprintf("%10s", "love")
	fmt.Println(z)

	n := rand.Intn(100)
	fmt.Println(n)
	n = rand.Intn(100)
	fmt.Println(n)

	// 将2进制转换为十进制， 二进制100000为32
	v, _ := strconv.ParseInt("100000", 2, 32)
	fmt.Println(v)
	// 十进制转换为2进制
	fmt.Println(strconv.FormatInt(32, 2))

	// Itoa 是 FormatInt(n, 2)的简写
	fmt.Println(strconv.Itoa(32))
	// 字符串转为int
	vv, _ := strconv.Atoi("32")
	fmt.Println(vv)

}
