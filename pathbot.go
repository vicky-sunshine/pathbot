package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type PathbotDirection struct {
	Direction string `json:"direction"`
}

type PathbotLocation struct {
	Status            string   `json:"status"`
	Message           string   `json:"message"`
	Exits             []string `json:"exits"`
	Description       string   `json:"description"`
	MazeExitDirection string   `json:"mazeExitDirection"`
	MazeExitDistance  int      `json:"mazeExitDistance"`
	LocationPath      string   `json:"locationPath"`
}

type RoomID struct {
	X int
	Y int
}

func (r RoomID) String() string {
	return fmt.Sprintf("(%v, %v)", r.X, r.Y)
}

// global visit map
var visit = make(map[RoomID]bool)

func main() {
	startLocation := start()
	startRoomID := RoomID{X: 0, Y: 0}

	DFSWalk(startLocation, startRoomID)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func DFSWalk(location PathbotLocation, roomID RoomID) (finished bool) {
	// print and save location info
	printLocation(location, roomID)
	visit[roomID] = true

	// check if game is finished
	if location.Status == "finished" {
		return true
	}

	// decide next direction
	printPrompt(location.Exits)
	for _, v := range location.Exits {
		var nextRoomID RoomID
		switch v {
		case "N":
			nextRoomID = RoomID{X: roomID.X + 1, Y: roomID.Y}
		case "S":
			nextRoomID = RoomID{X: roomID.X - 1, Y: roomID.Y}
		case "E":
			nextRoomID = RoomID{X: roomID.X, Y: roomID.Y + 1}
		case "W":
			nextRoomID = RoomID{X: roomID.X, Y: roomID.Y - 1}
		default:
			panic(fmt.Errorf("Un-possible direction %v", v))
		}

		if visited, _ := visit[nextRoomID]; !visited {
			printDecision(v)
			dir := PathbotDirection{Direction: v}
			body, err := json.Marshal(dir)
			checkError(err)

			nextLocation := apiPost(location.LocationPath, bytes.NewBuffer(body))
			if finished := DFSWalk(nextLocation, nextRoomID); finished {
				return finished
			}
		}
	}

	return false
}

// pathbot interact interface
func printPrompt(directions []string) {
	fmt.Println("What direction will you go?", directions)
}

func printDecision(direction string) {
	fmt.Println("Choose", direction)
}

func printLocation(location PathbotLocation, roomID RoomID) {
	fmt.Println()
	fmt.Println("[", location.Status, "|", "Room ID:", roomID, "]")
	fmt.Println(location.Message)
	fmt.Println(location.Description)
}

func start() PathbotLocation {
	return apiPost("/pathbot/start", strings.NewReader("{}"))
}

func apiPost(path string, body io.Reader) PathbotLocation {
	domain := "https://api.noopschallenge.com"
	res, err := http.Post(domain+path, "application/json", body)
	checkError(err)

	return parseResponse(res)
}

func parseResponse(res *http.Response) PathbotLocation {
	var response PathbotLocation
	body, err := ioutil.ReadAll(res.Body)
	checkError(err)

	err = json.Unmarshal(body, &response)
	checkError(err)

	return response
}
