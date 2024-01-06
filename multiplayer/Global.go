package multiplayer

import (
	"encoding/json"
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"net"
)

var Conn net.Conn = nil
var Map [][]int = [][]int{{}}
var MapReceived bool = false
var WaitingForResponse bool = false
var ServerPos map[string]int = map[string]int{"X": 0, "Y": 0}
var ClientPos map[string]int = map[string]int{"X": 0, "Y": 0}
var KeyPressed string = ""
var MultiplayerPortal [][]int = [][]int{}

func IsThereAPlayer(x, y, mapWidth, mapHeight int) bool {
	if !configuration.Global.TerreRonde {
		return (x == ServerPos["X"] && y == ServerPos["Y"]) || (x == ClientPos["X"] && y == ClientPos["Y"])
	} else {
		negativeX, negativeY := x, y
		if x < 0 {
			negativeX = mapWidth + x
		}
		if y < 0 {
			negativeY = mapHeight + y
		}
		return (x%mapWidth == ServerPos["X"] && y%mapHeight == ServerPos["Y"]) || (x%mapWidth == ClientPos["X"] && y%mapHeight == ClientPos["Y"]) || (negativeX == ServerPos["X"] && negativeY == ServerPos["Y"]) || (negativeX == ClientPos["X"] && negativeY == ClientPos["Y"])
	}
}

func SendMap() {
	if Conn != nil {
		JSONData := map[string]interface{}{
			"API":  "SendMap",
			"Data": Map,
		}
		data, _ := json.Marshal(JSONData)
		WaitingForResponse = true
		Conn.Write(data)
		for WaitingForResponse {
		}
		return
	}
}

func UpdateMap(data interface{}) [][]int {
	value2 := data.([]interface{})
	value := convertInterFaceArrayToArrayArrayInt(value2)
	return value
}

func convertInterFaceArrayToArrayArrayInt(array []interface{}) [][]int {
	var result [][]int
	for i := 0; i < len(array); i++ {
		var temp []int
		for x := 0; x < len(array[i].([]interface{})); x++ {
			temp = append(temp, int(((array[i].([]interface{}))[x]).(float64)))
		}
		result = append(result, temp)
	}
	return result
}

func SendPos(x, y int) {
	if Conn != nil {
		JSONData := map[string]interface{}{
			"API":  "SendPos",
			"Data": map[string]int{"X": x, "Y": y},
		}
		data, _ := json.Marshal(JSONData)
		WaitingForResponse = true
		Conn.Write(data)
		for WaitingForResponse {
		}
		return
	}
}

func SendKeyPressed(key string) {
	if Conn != nil {
		JSONData := map[string]interface{}{
			"API":  "SendKeyPressed",
			"Data": key,
		}
		data, _ := json.Marshal(JSONData)
		WaitingForResponse = true
		Conn.Write(data)
		for WaitingForResponse {
		}
		return
	}
}

func waitForResponse() {
	buffer := make([]byte, 1024)
	// Handle client connection in a goroutine
	bytesRead, err := Conn.Read(buffer)
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
	WaitingForResponse = false
	return
}

func datatReceived() {
	if Conn != nil {
		JSONData := map[string]interface{}{
			"API": "DataReceived",
		}
		data, _ := json.Marshal(JSONData)
		Conn.Write(data)
	}
}
