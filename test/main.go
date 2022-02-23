package main

import (
	"fmt"
	"strings"
)

func main() {
	split := strings.Split("tk01-test-log-containerd", "-")
	if split[0] == "tk01" {
		split = split[1:]
	}
	var stdoutContainerdLogIndex string
	if len(split) >= 3 {
		stdoutContainerdLogIndex = split[0] + "-containerd_stdout-" + split[1] + "-" + split[2]
	} else {
		stdoutContainerdLogIndex = "sscsys-containerd_stdout-nostd-nostd"
	}
	fmt.Println(split)
	fmt.Println(stdoutContainerdLogIndex)
}
