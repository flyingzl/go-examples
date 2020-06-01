package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"time"
)

func getTitle(url string) string {
	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	reg, _ := regexp.Compile(`(?i)<title.*?>(.*)</title>`)
	results := reg.FindStringSubmatch(string(body))
	if len(results) == 2 {
		return results[1]
	}
	return "unkown"

}

func worker(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {

	for {
		if url, more := <-jobs; more {
			title := getTitle(url)
			results <- fmt.Sprintf("%s[%s]", title, url)

		} else {
			wg.Done()
			return
		}
	}
}

func main() {

	threadsNum := 20

	jobs := make(chan string, threadsNum)

	// done := make(chan bool)
	urls := []string{"https://dict.youdao.com", "http://www.youku.com", "http://www.sina.com.cn", "https://www.zhihu.com"}

	wg := new(sync.WaitGroup)
	wg1 := new(sync.WaitGroup)
	wg.Add(threadsNum)
	wg1.Add(1)

	results := make(chan string, 10)

	// done := make(chan bool)

	for i := 1; i <= threadsNum; i++ {
		go worker(jobs, results, wg)
	}

	go func() {
		for v := range results {
			fmt.Println(v)
		}
		time.Sleep(2 * time.Second)
		fmt.Println("hello")
		wg1.Done()
	}()

	for _, url := range urls {
		jobs <- url
	}

	close(jobs)

	wg.Wait()

	close(results)

	wg1.Wait()

}
