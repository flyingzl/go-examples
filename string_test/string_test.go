package string_test

import (
	"fmt"
	"strconv"
	"testing"
	"unicode/utf8"
)

func TestStringConv(t *testing.T) {
	number := 10
	str := strconv.Itoa(number)
	t.Logf("%T", str)
	str = fmt.Sprintf("%s1", str)
	if number, err := strconv.Atoi(str); err == nil {
		t.Log(number)
	}
	t.Log(number)

}

func TestStringCount(t *testing.T) {
	greetings := "hello,石头"
	// bytes ---> rune
	greetingRune := []rune(greetings)
	// 字符串长度为12， 因为包含中文，一个中文长度为3
	t.Logf("unicodeLength=%d, byteLength=%d", len(greetings), len(greetingRune))
	// 返回可见的字符串长度，8
	t.Log(utf8.RuneCountInString(greetings))
}

func TestRuneToString(t *testing.T) {
	greetings := "hello,石头"
	// bytes ---> rune
	greetingRune := []rune(greetings)
	for _, v := range greetingRune {
		t.Logf("%[1]c, %[1]x", v)
	}

	// rune to string
	greetings = string(greetingRune[7:])
	t.Log(greetings)

	// string to byte
	t.Log([]byte(greetings))

	// byte to string
	z := []byte{229, 164, 180}
	t.Log(string(z))
}
