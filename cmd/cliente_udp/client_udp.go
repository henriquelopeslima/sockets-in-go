package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error){
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Uso: %s host:porta", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	addr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, addr)
	checkError(err)

	_, _ = conn.Write([]byte("Qualquer coisa"))
	var buf [512]byte
	n, _ := conn.Read(buf[0:])
	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}