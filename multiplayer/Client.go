package multiplayer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"net"
	"os"
)

func InitAsClient() {
	conn, err := net.Dial("tcp", configuration.Global.MultiplayerIP+":"+configuration.Global.ServerPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	serverResponse, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(serverResponse)
	if serverResponse != "validated\n" {
		fmt.Println(err)
		os.Exit(1)
	}
	Conn = conn
	fmt.Println("connection validated with " + conn.RemoteAddr().String())
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
			fmt.Println("Error3:", err)
			return
		}
		switch jsonData["API"] {
		case "SendMap":
			Map = UpdateMap(jsonData["Data"])
			datatReceived()
			fmt.Println("Map received")
		case "SendPos":
			ServerPos["X"] = int((jsonData["Data"].(map[string]interface{}))["X"].(float64))
			ServerPos["Y"] = int((jsonData["Data"].(map[string]interface{}))["Y"].(float64))
			MapReceived = true
			datatReceived()
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
