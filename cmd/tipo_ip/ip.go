package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Deve ser passado apenas o ip para ser validado\n")
		_, _ = fmt.Fprintf(os.Stderr, "Uso: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	//IPAddr
	addr := net.ParseIP(os.Args[1])

	if addr == nil {
		fmt.Println("Endereço inválido.")
	} else {
		fmt.Println("O endereço é: ", addr.String())
	}

	os.Exit(0)
}
