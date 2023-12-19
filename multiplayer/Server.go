package multiplayer

import (
	"encoding/json"
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"net"
)

func ConnectAsServer() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:"+configuration.Global.ServerPort)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listening on localhost:" + configuration.Global.ServerPort)
	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		// Handle client connection in a goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	fmt.Println("connection try")
	message := "validated\n"
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("connection validated with "+conn.RemoteAddr().String(), "on port "+configuration.Global.ClientPort)
	Conn = conn
	SendMap()
	buffer := make([]byte, 1024)
	for {
		// Handle client connection in a goroutine
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error1:", err)
			return
		}
		if bytesRead == 0 {
			fmt.Println("Error2: connection perdus")
			return
		}
		data := buffer[:bytesRead]
		jsonData := map[string]interface{}{}
		err = json.Unmarshal(data, &jsonData)
		if err != nil {
			fmt.Println(string(data))
			fmt.Println("Error3:", err)
			return
		}
		switch jsonData["API"] {
		case "InitReiceved":
			if jsonData["Data"].(bool) {
				SendPos(ServerPos["X"], ServerPos["Y"])
			} else {
				conn.Close()
			}
		}
		buffer = make([]byte, 1024)
	}
}
