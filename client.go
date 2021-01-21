package main

import (
	"net"
 	"fmt"
 	"bufio"
	"os"
)

func main() {
	
	client, error_1 := net.Dial("tcp", "127.0.0.1:8081")

	if error_1 != nil {
		fmt.Println(error_1)
		os.Exit(3)
	}

	for {

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Text to sent: ")
		text, error_2 := reader.ReadString('\n')

		if error_2 != nil {
			fmt.Println(error_2)
			os.Exit(3)
		}

		fmt.Fprintf(client, text + "\n")

		message, error_3 := bufio.NewReader(client).ReadString('\n')

		if error_3 != nil {
			fmt.Println(error_3)
			os.Exit(3)
		}

		fmt.Print("Server Response: " + message)

	}
}
