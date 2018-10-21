package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://www.google.com.hk:443/search?q=love&lan=zh_CN#a=3&b=4"
	u, _ := url.Parse(s)
	if u != nil {
		fmt.Println(u.Scheme)
		fmt.Println(u.Host)
		fmt.Println(u.Path)
		fmt.Println(u.Fragment)
		fmt.Println(u.RawQuery)
		fmt.Println(u.Hostname())

		m, _ := url.ParseQuery(u.RawQuery)
		fmt.Println(m["q"][0])
	}
}
