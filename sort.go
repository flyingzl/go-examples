// 在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。
// 对于更加复杂的情况，我们可以使用一个_[互斥锁](http://zh.wikipedia.org/wiki/%E4%BA%92%E6%96%A5%E9%94%81)_
// 来在 Go 协程间安全的访问数据。

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// ByLength 用于比较字符串的长度
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("Point<x:%d, y:%d>\n", p.x, p.y)
}

func main() {
	a := []int{3, 0, 2, 1}
	sort.Ints(a)
	fmt.Println(a)
	fmt.Println(sort.IntsAreSorted(a))

	b := []string{"2", "1a", "a"}
	sort.Strings(b)
	fmt.Println(b)
	fmt.Println(sort.StringsAreSorted(b))

	str := []string{"a1", "bccc", "2"}
	sort.Sort(ByLength(str))
	fmt.Println(str)

	f, err := os.Create("c:/demo.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString(strings.ToUpper("hello"))

	fmt.Println(len([]rune("你好")))

	fmt.Printf("|%6d|%6d|\n", 100, 200)
	fmt.Printf("|%-6d|%-6d|\n", 100, 200)
	fmt.Println(fmt.Sprintf("|hello %6s|\n", "love"))
	fmt.Println(fmt.Sprintf("|hello %-6s|\n", "love"))

	aa := Point{x: 2, y: 3}
	fmt.Println(&aa)
	fmt.Printf("%q", "hello")

}
