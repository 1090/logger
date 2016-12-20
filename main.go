package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.0.4:30003")
	//conn, err := net.Dial("tcp", "127.0.0.1:30003")

	if err != nil {
		// handle error
		fmt.Printf("connection error: %s\n", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(conn)

	for {
		output, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		if len(strings.TrimSpace(output)) != 0 {
			fmt.Print(output)
		}
	}

}
