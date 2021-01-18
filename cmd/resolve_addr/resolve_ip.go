package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Uso: %s hostname\n", os.Args[0])
		os.Exit(1)
	}

	addr, err := net.ResolveIPAddr("ip6", os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Endereço resolvido: ", addr.String())

	//slice = array de tamanho variável
	addrs, err := net.LookupHost(os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Endereços: ")
	for _, s := range addrs {
		fmt.Println(s)
	}

	os.Exit(0)
}
