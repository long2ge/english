package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// cmd := exec.Command("open -a \"Google Chrome\" http://baidu.com")
	// test := cmd.Run()

	// fmt.Println(test)

	cmd := exec.Command("open", "-a", "Google Chrome", "https://baidu.com")

	cmd.Run()

	// cmd1 := exec.Command("ls")

	// cmd1.Run()

	fmt.Println(cmd.Run())
}
