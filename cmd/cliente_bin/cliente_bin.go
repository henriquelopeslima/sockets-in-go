package main

import (
	"bytes"
	"encoding/binary"
	"net"
	"os"
	"time"
)

type msg struct {
	SensID uint16
	LocID uint16
	Tstamp uint32
	Temp int16
}

func main() {

	data := msg{
		SensID: 1,
		LocID: 234,
		Tstamp: uint32(time.Now().Unix()),
		Temp: 12,
	}

	buffer := new(bytes.Buffer)

	_ = binary.Write(buffer, binary.BigEndian, data)

	service := ":1200"
	addr, _ := net.ResolveUDPAddr("udp", service)
	conn, _ := net.DialUDP("udp", nil, addr)

	_, _ = conn.Write(buffer.Bytes())

	os.Exit(0)
}
