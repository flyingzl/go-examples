package main

import (
	"time"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	s := "admin"
	hash := md5.New()
	hash.Write([]byte(s))

	bs := hash.Sum(nil)
	value := fmt.Sprintf("%x", bs)
	fmt.Println(value)

	mac := hmac.New(sha256.New, []byte("hello"))
	mac.Write([]byte("love"))
	bs = mac.Sum(nil)
	fmt.Printf("%x\n", bs)

	s1 := btoa([]byte("hello"))
	fmt.Println(s1)

	b1, _ := atob(s1)
	fmt.Println(string(b1))

}

func btoa(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func atob(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
