package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func StartServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		fmt.Print("Enter command: ")
		command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		command = strings.TrimSpace(command)
		if command == "exit" {
			fmt.Println("Exiting...")
			return
		}
		conn.Write([]byte(command + "\n"))

		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			return
		}
		fmt.Println("Response from client:", response)
	}
}
