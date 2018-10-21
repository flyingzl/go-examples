package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		v := strings.ToUpper(scanner.Text())
		if v == "Q" {
			os.Exit(1)
		}
		fmt.Println(v)
	}
}
