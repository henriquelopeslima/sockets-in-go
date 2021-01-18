package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError(err error){
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte

	_, addr, _ := conn.ReadFromUDP(buf[0:])

	dayTime := time.Now().String()

	_, _ = conn.WriteToUDP([]byte(dayTime), addr)
}

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
}
