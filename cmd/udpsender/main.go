package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"fmt"
)

func main (){
	addr, err := net.ResolveUDPAddr("udp", ":42069")

	if err != nil {
		log.Fatal(err)
	}

	conn, connErr := net.DialUDP("udp", nil, addr) 
	
	if connErr != nil {
		log.Fatal(connErr)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(">")
		chunk := make([]byte, 8, 8)
		_, readErr := reader.Read(chunk)

		if readErr != nil {
			formattedErr := fmt.Errorf("%w", readErr)
			fmt.Println(formattedErr)
		}


		_, writeErr := conn.Write(chunk)
		if writeErr != nil {
			formattedErr := fmt.Errorf("%w", writeErr)
			fmt.Println(formattedErr)
		}

	}
	
}
