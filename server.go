package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){
	fmt.Println("Server waitting for clients...")

	server, error_1 := net.Listen("tcp", ":8081")

	if error_1 != nil {
		fmt.Println(error_1)
		os.Exit(3)
	}

	client, error_2 := server.Accept()

	if error_2 != nil {
		fmt.Println(error_2)
		os.Exit(3)
	}

	defer server.Close()

	fmt.Println("Connection accepted...")

	for {

		message, error_3 := bufio.NewReader(client).ReadString('\n')

		if error_3 != nil {
			fmt.Println("An error has been found, client connection has gonna way.")
			os.Exit(3)
		}

		fmt.Print("Message has been received:" + string(message))

		formatted_message := strings.ToUpper(message)

		client.Write([]byte(formatted_message + "\n"))

	}
}
