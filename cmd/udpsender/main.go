package udpsender

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"fmt"
)

func main (){
	addr, err := net.ResolveUDPAddr("UDP", ":42069")

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
		i, readErr := reader.Read(chunk)

		if readErr != nil {
			formattedErr := fmt.Errorf("%w", readErr)
			fmt.Println(formattedErr)
		}



	}
	
}
