package main

import (
	"fmt"
	"strings"
	"io"
	"log"
	"net"
	"errors"
)

func main() {
	listener, err := net.Listen("tcp", ":42069")
	defer listener.Close() 
	
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil  {
			log.Fatal(err)
		}

		fmt.Println("Accepted connection from: ", conn.RemoteAddr())

		channel := getLinesChannel(conn)
		for message := range channel {
			fmt.Println(message)

		}
	}
	
}


func getLinesChannel(f io.ReadCloser) <-chan string {
	channel := make(chan string)
	chunk := make([]byte, 8, 8)

	currentLine := ""
	go func (){
		for {
			_, err := f.Read(chunk)
			if err != nil {
				if currentLine != "" {
					channel <- currentLine
				}
				if errors.Is(err, io.EOF) {
					break
				} else {
					fmt.Println("the connection has been closed")
					close(channel)
				}
			}

			splits := strings.Split(string(chunk), "\n")
			currentLine += splits[0]

			if len(splits) > 1 {
				channel <- currentLine
				currentLine = splits[1]
			}
		}
	}()

	return channel

}
