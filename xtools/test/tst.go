package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Println("niubi")
		for _, interf := range interfaces {
			fmt.Printf("\r\ninterf:%+v\r\n", interf)
		}
	}
}
