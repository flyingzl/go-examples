package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	lsCmd := exec.Command("ls", "-l")
	output, _ := lsCmd.Output()
	fmt.Println("> ls -l | wc -l")

	wcCmd := exec.Command("wc", "-l")
	wcIn, _ := wcCmd.StdinPipe()
	wcOut, _ := wcCmd.StdoutPipe()
	wcCmd.Start()
	wcIn.Write(output)
	wcIn.Close()
	wcBytes, _ := ioutil.ReadAll(wcOut)
	wcCmd.Wait()
	fmt.Println(string(wcBytes))

	os.Setenv("PATH", "/Users/zhaolei/go/bin"+":"+os.Getenv("PATH"))
	binary, _ := exec.LookPath("node")
	err := syscall.Exec(binary, []string{"node", "-v"}, nil)
	if err != nil {
		panic(err)
	}

}
