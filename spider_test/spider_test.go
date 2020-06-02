package spider_test

import (
	"github.com/gocolly/colly"
	"testing"
)

func TestSpider(t *testing.T) {
	c := colly.NewCollector()
	// Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		t.Log(e.Request.AbsoluteURL(e.Attr("href")))
	})
	c.OnRequest(func(r *colly.Request) {
		t.Log("Visiting", r.URL)
	})

	c.Visit("https://www.baidu.cn")
}
