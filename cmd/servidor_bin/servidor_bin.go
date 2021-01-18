package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type msg struct {
	SensID uint16
	LocID uint16
	Tstamp uint32
	Temp int16
}

func handleClient(conn *net.UDPConn)  {
	var buffer = make([]byte, 10)
	var data msg

	_, _, _ = conn.ReadFromUDP(buffer)

	_ = binary.Read(bytes.NewReader(buffer), binary.BigEndian, &data)

	fmt.Println("Sensor: ", data.SensID)
	fmt.Println("Localização: ", data.LocID)
	fmt.Println("Timestamp: ", data.Tstamp)
	fmt.Println("Temperatura: ", data.Temp)
}

func main() {
	service := ":1200"
	addr, _ := net.ResolveUDPAddr("udp", service)

	conn, _ := net.ListenUDP("udp", addr)

	for {
		handleClient(conn)
	}
}
