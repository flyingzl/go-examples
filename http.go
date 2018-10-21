package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// IPInfo 用来模拟IP数据对象
type IPInfo struct {
	IP      string `json:"ip"`
	Region  string `json:"region"`
	Country string `json:"country"`
	Org     string `json:"org"`
}

func (v IpInfo) String() string {
	return fmt.Sprintf("<IPInfo>(IP=%s, Region=%s, Country=%s, Org=%s)\n", v.IP, v.Region, v.Country, v.Org)
}

func main() {
	url := "https://ipinfo.io/"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// 通过ioutill读取返回的数据对象
	body, _ := ioutil.ReadAll(res.Body)
	ipInfo := &IPInfo{}
	if err := json.Unmarshal(body, ipInfo); err != nil {
		panic(err)
	}
	fmt.Println(ipInfo)

}
