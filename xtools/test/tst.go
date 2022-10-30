package main

import (
	"fmt"
	"runtime"
)

func main() {

	os := runtime.GOOS
	fmt.Printf("os:%+v\n", os)

}
