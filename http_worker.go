package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getTitle(url string) string {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	reg, _ := regexp.Compile(`(?i)<title>(.*)</title>`)
	results := reg.FindStringSubmatch(string(body))
	if len(results) == 2 {
		return results[1]
	}
	return "unkown"

}

func worker(id int, jobs <-chan string, results chan<- string) {
	for url := range jobs {
		// fmt.Println("worker", id, "processing url ", url)
		title := getTitle(url)
		results <- fmt.Sprintf("%s[%s]", title, url)
	}
}

func main() {
	jobs := make(chan string, 10)
	results := make(chan string, 10)
	urls := []string{
		"http://www.sogou.com",
		"http://www.baidu.com",
		"http://www.migu.cn",
		"https://asyncoder.com",
		"http://www.taobao.com",
		"http://www.sina.com",
		"http://www.youku.com",
		"http://www.360.cn",
		"http://m.163.com",
		"http://www.toutiao.com",
	}
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	for i := 0; i < len(urls); i++ {
		v := <-results
		fmt.Println(v)
	}

}
