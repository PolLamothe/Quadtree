package multiplayer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
)

var Conn net.Conn = nil
var Map [][]int = [][]int{{}}
var MapReceived bool = false
var WaitingForResponse bool = false
var ServerPos map[string]int = map[string]int{"X": 0, "Y": 0}
var ClientPos map[string]int = map[string]int{"X": 0, "Y": 0}
var KeyPressed string = ""
var MultiplayerPortal [][]int = [][]int{}
var BlockToSend []map[string]int = []map[string]int{}
var ReceivingBlock bool = false
var RoutineFinished bool = false
var SendingConfirmation bool = false

func StoreInFile(x, y, value int) {
	f, err := os.OpenFile("../multiplayer/BlockGeneratedServer", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	data, _ := json.Marshal(map[string]int{"X": x, "Y": y, "Value": value})
	if _, err := f.WriteString(string(data) + "\n"); err != nil {
		log.Println(err)
	}
}

func SendPortal() {
	if Conn != nil {
		JSONData := map[string]interface{}{
			"API":  "SendPortal",
			"Data": MultiplayerPortal,
		}
		data, _ := json.Marshal(JSONData)
		WaitingForResponse = true
		for SendingConfirmation {
			time.Sleep(100 * time.Millisecond)
		}
		Conn.Write(data)
		for WaitingForResponse {
		}
		MultiplayerPortal = [][]int{}
	}
}

func SendConfig() {
	if Conn != nil {
		ConfigDATA := map[string]interface{}{
			"RandomGeneration":    configuration.Global.RandomGeneration,
			"RandomTileX":         configuration.Global.RandomTileX,
			"RandomTileY":         configuration.Global.RandomTileY,
			"Portal":              configuration.Global.Portal,
			"SingleUsagePortal":   configuration.Global.SingleUsagePortal,
			"CameraBlockEdge":     configuration.Global.CameraBlockEdge,
			"CameraFluide":        configuration.Global.CameraFluide,
			"GenerationInfinie":   configuration.Global.GenerationInfinie,
			"TerreRonde":          configuration.Global.TerreRonde,
			"MultiplayerColision": configuration.Global.MultiplayerColision,
		}
		JSONData := map[string]interface{}{
			"API":  "SendConfig",
			"Data": ConfigDATA,
		}
		data, _ := json.Marshal(JSONData)
		WaitingForResponse = true
		for SendingConfirmation {
			time.Sleep(100 * time.Millisecond)
		}
		Conn.Write(data)
		for WaitingForResponse {
		}
		fmt.Println("config sent succesfully")
		return
	}
}

func StartSendingBlock() {
	if Conn != nil {
		JSONData := map[string]interface{}{
			"API": "StartSendingBlock",
		}
		data, _ := json.Marshal(JSONData)
		WaitingForResponse = true
		for SendingConfirmation {
			time.Sleep(100 * time.Millisecond)
		}
		Conn.Write(data)
		for WaitingForResponse {
		}
	}
}

func IsThisBlockReceived(x, y int) (bool, int) {
	//Ouverture du fichier
	var path string
	if configuration.Global.MultiplayerKind == 1 {
		path = "../multiplayer/BlockGeneratedServer"
	} else {
		path = "../multiplayer/BlockGeneratedClient"
	}
	path, err := filepath.Abs(path)
	if err != nil {
		os.Exit(1)
	}
	myFile, err2 := os.Open(path)
	if err2 != nil {
		log.Print(err2)
		os.Exit(1)
	}
	//Préparation de la lecture
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(myFile)
	// Lecture des lignes du fichier
	var tempFile, err4 = os.Create("../multiplayer/temp")
	if err4 != nil {
		log.Fatal(err4)
		os.Exit(1)
	}
	defer tempFile.Close()
	var found bool = false
	var value int = 0
	for scanner.Scan() {
		var current string = scanner.Text()
		var data map[string]int = map[string]int{}
		err := json.Unmarshal([]byte(current), &data)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if data["X"] == x && data["Y"] == y {
			found = true
			value = data["Value"]
		} else {
			if _, err := tempFile.Write(append([]byte(current), '\n')); err != nil {
				log.Println(err)
				os.Exit(1)
			}
		}
	}
	if found {
		os.RemoveAll(path)
		var newPath string
		if configuration.Global.MultiplayerKind == 1 {
			newPath = "../multiplayer/BlockGeneratedServer"
		} else {
			newPath = "../multiplayer/BlockGeneratedClient"
		}
		os.Rename("../multiplayer/temp", newPath)
	}
	// Fermeture du fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	return found, value
}

func treatBlocReceived(jsonData map[string]interface{}) {
	temp := jsonData["Data"].([]interface{})
	var temp2 []map[string]int
	for v := range temp {
		temp3 := temp[v].(map[string]interface{})
		var temp4 map[string]int = map[string]int{"X": int(temp3["X"].(float64)), "Y": int(temp3["Y"].(float64)), "Value": int(temp3["Value"].(float64))}
		temp2 = append(temp2, temp4)
		data, err := json.Marshal(temp4)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		data = append(data, []byte("\n")...)
		var path string
		if configuration.Global.MultiplayerKind == 2 {
			path, _ = filepath.Abs("../multiplayer/BlockGeneratedClient")
		} else {
			path, _ = filepath.Abs("../multiplayer/BlockGeneratedServer")
		}
		f, err2 := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err2 != nil {
			log.Print(err2)
			os.Exit(1)
		}
		defer f.Close()
		if _, err := f.Write(data); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
	DatatReceived()
}

func StopSendingBlock() {
	if Conn != nil {
		JSONData := map[string]interface{}{
			"API": "StopSendingBlock",
		}
		data, _ := json.Marshal(JSONData)
		WaitingForResponse = true
		for SendingConfirmation {
			time.Sleep(100 * time.Millisecond)
		}
		Conn.Write(data)
		for WaitingForResponse {
		}
	}
}

func SendBlock() {
	if Conn != nil {
		StartSendingBlock()
		if RoutineFinished {
			for len(BlockToSend) > 0 { //tant qu'il reste des blocs a envoyer on va les envoyers par paquet de 10 (sinon il y'en a trop et la fonction unmarshall ne fonctionne pas)
				var temp []map[string]int = []map[string]int{}
				for x := 0; x < 10 && len(BlockToSend) > 0; x++ {
					temp = append(temp, BlockToSend[0])
					BlockToSend = BlockToSend[1:]
				}
				JSONData := map[string]interface{}{
					"API":  "SendBlock",
					"Data": temp,
				}
				data, _ := json.Marshal(JSONData)
				WaitingForResponse = true
				for SendingConfirmation {
					time.Sleep(100 * time.Millisecond)
				}
				Conn.Write(data)
				for WaitingForResponse {
				}
			}
		} else {
			//Ouverture du fichier
			var filePath string = "../multiplayer/BlockGeneratedServer"
			var myFile *os.File
			var err error
			myFile, err = os.Open(filePath)
			//Préparation de la lecture
			var scanner *bufio.Scanner
			scanner = bufio.NewScanner(myFile)
			// Lecture des lignes du fichier
			var temp []map[string]int
			for scanner.Scan() {
				var data map[string]int = map[string]int{}
				err := json.Unmarshal([]byte(scanner.Text()), &data)
				if err != nil {
					log.Fatal(err)
				}
				temp = append(temp, data)
				if len(temp) == 10 {
					JSONData := map[string]interface{}{
						"API":  "SendBlock",
						"Data": temp,
					}
					data, _ := json.Marshal(JSONData)
					WaitingForResponse = true
					for SendingConfirmation {
						time.Sleep(100 * time.Millisecond)
					}
					Conn.Write(data)
					for WaitingForResponse {
					}
					temp = []map[string]int{}
				}
			}
			if len(temp) > 0 {
				JSONData := map[string]interface{}{
					"API":  "SendBlock",
					"Data": temp,
				}
				data, _ := json.Marshal(JSONData)
				WaitingForResponse = true
				for SendingConfirmation {
					time.Sleep(100 * time.Millisecond)
				}
				Conn.Write(data)
				for WaitingForResponse {
				}
			}
			// Fermeture du fichier
			err = myFile.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
		RoutineFinished = true
		StopSendingBlock() //on prévient l'autre que l'on a fini d'envoyer les blocs
	}
}

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
		for SendingConfirmation {
			time.Sleep(100 * time.Millisecond)
		}
		Conn.Write(data)
		for WaitingForResponse {
		}
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
		for SendingConfirmation {
			time.Sleep(100 * time.Millisecond)
		}
		Conn.Write(data)
		for WaitingForResponse {
		}
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
		for SendingConfirmation {
			time.Sleep(100 * time.Millisecond)
		}
		Conn.Write(data)
		for WaitingForResponse {
		}
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

func DatatReceived() {
	if Conn != nil {
		SendingConfirmation = true
		JSONData := map[string]interface{}{
			"API": "DataReceived",
		}
		data, _ := json.Marshal(JSONData)
		Conn.Write(data)
		SendingConfirmation = false
	}
}
