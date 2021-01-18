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

func handleClient(conn net.Conn)  {
	defer conn.Close()

	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}

		fmt.Println(string(buf[0:]))

		_, err2 := conn.Write(buf[0:n])

		if err2 != nil {
			return
		}
	}
}

func main() {
	service := ":1220"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}