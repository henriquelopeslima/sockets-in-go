package main

import (
	"fmt"
	"io/ioutil"
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

	//TCPAddr
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	//TCPConn
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)

	fmt.Println(string(result))

	os.Exit(0)
}
