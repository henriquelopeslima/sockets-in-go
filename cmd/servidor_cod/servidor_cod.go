package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
)

func handleClient(conn *net.UDPConn) {
	var netBuffer [8000]byte
	var fileBuffer [8000]byte

	_, _, _ = conn.ReadFromUDP(netBuffer[0:])

	_, _ = base64.StdEncoding.Decode(fileBuffer[:], netBuffer[:])

	_ = ioutil.WriteFile("go.jpeg", fileBuffer[:], 0444)
}

func handleClientText(conn *net.UDPConn) {
	var buf [512]byte
	_, _, _ = conn.ReadFromUDP(buf[0:])
	fmt.Println("Dado codificado ", string(buf[:]))
	dec, _ := base64.StdEncoding.DecodeString(string(buf[:]))
	fmt.Println("Dado decodificado ", dec)
}

func main() {
	service := ":1200"
	addr, _ := net.ResolveUDPAddr("udp", service)

	conn, _ := net.ListenUDP("udp", addr)

	for {
		handleClient(conn)
	}

}
