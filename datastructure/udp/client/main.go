package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr , err := net.ResolveUDPAddr("udp", "0.0.0.0:9998")
	if err != nil {
		log.Println(err)
	}
	conn, err := net.DialUDP("udp", nil,addr)
	if err != nil {
		log.Println(err)
	}
	_, err = conn.Write([]byte("日志信息"))
	if err != nil {
		log.Println(err)
	}
	var buf = make([]byte, 1024)
	_, addr, _ = conn.ReadFromUDP(buf)
	fmt.Println("accept server message", string(buf))
}