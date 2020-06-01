package package_test

import (
	"fmt"
	cowsay "github.com/Code-Hex/Neo-cowsay"
	"go-examples/hello"
	"testing"
)

func init() {
	fmt.Println("test init..")
}

func TestPackageInit(t *testing.T) {
	hello.Greet("hello")
	t.Log("invoked...")
	s, err := cowsay.Say(cowsay.Phrase("真是美好的一天"), cowsay.Type("default"), cowsay.Rainbow())
	if err == nil {
		t.Log(s)
	}
}
