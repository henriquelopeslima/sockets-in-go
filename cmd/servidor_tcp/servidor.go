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

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1200")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		fmt.Println("Conectado")
		if err != nil {
			continue
		}

		dayTime := time.Now().String()
		_, _ = conn.Write([]byte(dayTime))
		_ = conn.Close()
	}
}
