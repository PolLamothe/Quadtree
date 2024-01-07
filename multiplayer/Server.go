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
	go SendMap()
	waitForResponse()
	go SendPos(ServerPos["X"], ServerPos["Y"])
	waitForResponse()
	go SendBlock()
	waitForResponse()
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
		case "SendKeyPressed":
			KeyPressed = jsonData["Data"].(string)
			datatReceived()
		case "StartSendingBlock":
			ReceivingBlock = true
			datatReceived()
		case "StopSendingBlock":
			ReceivingBlock = false
			datatReceived()
		case "SendBlock":
			temp := jsonData["Data"].([]interface{})
			var temp2 []map[string]int
			for v := range temp {
				temp3 := temp[v].(map[string]interface{})
				var temp4 map[string]int = map[string]int{"X": int(temp3["X"].(float64)), "Y": int(temp3["Y"].(float64)), "Value": int(temp3["Value"].(float64))}
				temp2 = append(temp2, temp4)
			}
			BlockReceived = append(BlockReceived, temp2...)
			datatReceived()
		case "DataReceived":
			WaitingForResponse = false
		}
		buffer = make([]byte, 1024)
	}
}
