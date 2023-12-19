package multiplayer

import (
	"encoding/json"
	"net"
)

var Conn net.Conn
var Map [][]int = [][]int{{}}
var MapReceived bool = false
var ServerPos map[string]int = map[string]int{}
var ClientPos map[string]int = map[string]int{}

func SendMap() {
	JSONData := map[string]interface{}{
		"API":  "SendMap",
		"Data": Map,
	}
	data, _ := json.Marshal(JSONData)
	Conn.Write(data)
}

func InitReiceved() {
	JSONData := map[string]interface{}{
		"API":  "InitReiceved",
		"Data": true,
	}
	data, _ := json.Marshal(JSONData)
	Conn.Write(data)
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
	JSONData := map[string]interface{}{
		"API":  "SendPos",
		"Data": map[string]int{"X": x, "Y": y},
	}
	data, _ := json.Marshal(JSONData)
	Conn.Write(data)
}
