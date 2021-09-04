package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr , err := net.ResolveUDPAddr("udp", "0.0.0.0:9998")
	if err != nil {
		log.Fatal(err)
	}
	listener,err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for{
		var buf = make([]byte,1024)
		messagelen, uaddr, err := listener.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("read from client error: %s", err.Error())
			continue
		}
		fmt.Printf("message len is %d\n", messagelen)
		fmt.Printf("remote client connection is %s\n" , uaddr.String() )
		listener.WriteToUDP([]byte("ok"), uaddr)
	}
}
