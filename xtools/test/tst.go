package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func ExecCmd(c string, args ...string) string {
	cc := fmt.Sprintf("%s", args)
	cc1 := fmt.Sprintf("%s", strings.Split(cc, "[")[1])
	cc2 := fmt.Sprintf("%s", strings.Split(cc1, "]")[0])
	fmt.Printf("%+v\n", cc2)
	log.Printf("exec %v %v", c, cc2)
	cmd := exec.Command(c, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Println("failed to exec cmd:", err)
	}
	if len(out) == 0 {
		return ""
	}
	s := string(out)
	return strings.ReplaceAll(s, "\n", "")
}

func main() {
	ExecCmd("/usr/bin/ls", "aaa", "bbb", "ccc")
}
