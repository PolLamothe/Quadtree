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
	RoutineFinished = true
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
			fmt.Println(string(data))
			return
		}
		switch jsonData["API"] {
		case "SendMap":
			Map = UpdateMap(jsonData["Data"])
			DatatReceived()
			fmt.Println("Map received")
		case "SendPos":
			ServerPos["X"] = int((jsonData["Data"].(map[string]interface{}))["X"].(float64))
			ServerPos["Y"] = int((jsonData["Data"].(map[string]interface{}))["Y"].(float64))
			MapReceived = true
			DatatReceived()
		case "SendKeyPressed":
			KeyPressed = jsonData["Data"].(string)
			DatatReceived()
		case "StartSendingBlock":
			ReceivingBlock = true
			DatatReceived()
		case "StopSendingBlock":
			ReceivingBlock = false
			DatatReceived()
		case "SendBlock":
			treatBlocReceived(jsonData)
		case "SendConfig":
			var NewConfig map[string]interface{} = jsonData["Data"].(map[string]interface{})
			configuration.Global.RandomGeneration = NewConfig["RandomGeneration"].(bool)
			configuration.Global.RandomTileX = int(NewConfig["RandomTileX"].(float64))
			configuration.Global.RandomTileY = int(NewConfig["RandomTileY"].(float64))
			configuration.Global.Portal = NewConfig["Portal"].(bool)
			configuration.Global.SingleUsagePortal = NewConfig["SingleUsagePortal"].(bool)
			configuration.Global.CameraBlockEdge = NewConfig["CameraBlockEdge"].(bool)
			configuration.Global.CameraFluide = NewConfig["CameraFluide"].(bool)
			configuration.Global.GenerationInfinie = NewConfig["GenerationInfinie"].(bool)
			configuration.Global.TerreRonde = NewConfig["TerreRonde"].(bool)
			configuration.Global.MultiplayerColision = NewConfig["MultiplayerColision"].(bool)
			DatatReceived()
		case "DataReceived":
			WaitingForResponse = false
		}
		buffer = make([]byte, 1024)
	}
}
