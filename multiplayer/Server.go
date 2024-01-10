package multiplayer

import (
	"encoding/json"
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"net"
	"strings"
)

func ConnectAsServer() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", ":"+configuration.Global.ServerPort)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listening on localhost:" + configuration.Global.ServerPort)
	for {
		fmt.Println("connection received")
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
	fmt.Println("connection validated with " + conn.RemoteAddr().String())
	Conn = conn
	SendConfig()
	SendMap()
	SendPos(ServerPos["X"], ServerPos["Y"])
	SendBlock()
	SendPortal()
	fmt.Println("all data sent")
	RoutineFinished = true
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
		var dataString string = string(data)
		var dataArray []string = strings.SplitAfter(dataString, "}{")
		if len(dataArray) > 1 {
			dataArray[0] = dataArray[0][:len(dataArray[0])-1]
		}
		for i := 1; i < len(dataArray); i++ {
			dataArray[i] = "{" + dataArray[i]
		}
		for i := 0; i < len(dataArray); i++ {
			err = json.Unmarshal([]byte(dataArray[i]), &jsonData)
			if err != nil {
				fmt.Println(dataArray)
				fmt.Println("Error31:", err)
				return
			}
			switch jsonData["API"] {
			case "SendKeyPressed":
				KeyPressed = jsonData["Data"].(string)
				DatatReceived()
			case "SendBlock":
				treatBlocReceived(jsonData)
			case "DataReceived":
				WaitingForResponse = false
			}
		}
		buffer = make([]byte, 1024)
	}
}
