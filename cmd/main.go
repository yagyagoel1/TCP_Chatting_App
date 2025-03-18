package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
)

var connection = make(map[net.Conn]int)
var mutex = sync.Mutex{}

func decryptCommand(reader *bufio.Reader) (string, error) {
	//broadcast
	//send
	var command []string
	fmt.Println("reading")
	for {
		char, err := reader.ReadByte()
		if err != nil {
			fmt.Println("erorr while reading the command ")
			return "", err
		}
		fmt.Println(string(char), "char")
		if string(char) == "\n" {
			break
		}
		command = append(command, string(char))

	}
	return strings.Join(command, ""), nil
}
func getTheConnectionNumber(reader *bufio.Reader) (int, error) {
	var user []string
	for {
		char, err := reader.ReadByte()
		if err != nil {
			fmt.Println("erorr while reading the command ")
			return 0, err
		}

		if string(char) == "\n" {

			break

		}
		user = append(user, string(char))

	}
	userNum, err := strconv.Atoi(strings.Join(user, ""))
	if err != nil {
		return 0, err
	}
	return userNum, nil
}

func handleClient(conn net.Conn) {
	defer func() {
		mutex.Lock()
		delete(connection, conn)
		conn.Close()
		mutex.Unlock()
		fmt.Println("client disconnected ", conn)
	}()

	fmt.Println("connection established to client", conn)
	reader := bufio.NewReader(conn)
	for {
		command, err := decryptCommand(reader)
		if err != nil {
			fmt.Println("error while reading message from", conn, "error:", err)
			return
		}

		var user int
		if command == "send" {
			user, err = getTheConnectionNumber(reader)
			if err != nil {
				fmt.Println("error while reading message from", conn, "error:", err)
				return
			}
		}

		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error while reading message from", conn, "error:", err)
			return
		}

		fmt.Println("message received:", message)

		if command == "broadcast" {
			broadcastMessage(conn, message)
		} else if command == "send" {
			if err := sendMessage(message, user); err != nil {
				fmt.Println("Failed to send message:", err)

			}
		} else {
			fmt.Println("Unknown command:", command)
		}
	}
}
func sendMessage(message string, user int) error {
	mutex.Lock()
	defer mutex.Unlock()

	var conn net.Conn
	found := false

	for key, value := range connection {
		if value == user {
			conn = key
			found = true
			break
		}
	}

	if !found || conn == nil {
		return fmt.Errorf("user %d not found or connection is nil", user)
	}

	_, err := conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("problem while sending the message to the user: %w", err)
	}

	return nil
}
func broadcastMessage(sender net.Conn, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	for conn := range connection {
		if conn != sender {
			_, err := conn.Write([]byte("Message received from: " + strconv.Itoa(connection[sender]) + " | Message: " + message + "\n"))
			if err != nil {
				fmt.Println("error while writing to connection", conn, "error:", err)
				conn.Close()
				delete(connection, conn)
			}

		}
	}
}

func main() {
	//open a tcp server
	listner, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("error while starting the tcp server", err)
	}
	defer listner.Close()

	fmt.Println("Server is listening on port 8080")
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("error while acceptiing the connection ", err)
			continue

		}
		go handleClient(conn)

		connectionInt := rand.Intn(500000-100000) + 100000

		connection[conn] = connectionInt
		broadcastMessage(conn, "client connected ")
	}
}
