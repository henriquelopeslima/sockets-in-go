package main

import (
	"bufio"
	"encoding/base64"
	"net"
	"os"
)

func readImage() []byte {
	file, _ := os.Open("go.jpeg")
	defer file.Close()

	fileStat, _ := file.Stat()
	var fileSize = fileStat.Size()
	buffer := make([]byte, fileSize)

	reader := bufio.NewReader(file)
	_, _ = reader.Read(buffer)

	return buffer
}

func main() {

	//binary := []byte{1,2,3,4,5,6,7,8,9,10}b
	//Ou envie
	binary := readImage()

	enc := base64.StdEncoding.EncodeToString(binary)

	service := ":1200"
	addr, _ := net.ResolveUDPAddr("udp", service)
	conn, _ := net.DialUDP("udp", nil, addr)

	_, _ = conn.Write([]byte(enc))

	os.Exit(0)
}
