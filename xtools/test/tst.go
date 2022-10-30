package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println(net.IPv4(8, 8, 8, 8))
	fmt.Printf("%T\n", net.IPv4(8, 8, 8, 8))
}
