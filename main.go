package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type EventJumpOption struct {
	p      float32
	jumpID int
}

type Event struct {
	id       int
	name     string
	jumpList []EventJumpOption
}

func main() {
	// Read Event File
	file, err := os.Open("events")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer file.Close()

	// Parse the Event File
	scanner := bufio.NewScanner(file)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	var records [][]string
	for scanner.Scan() {
		records = append(records, strings.Split(scanner.Text(), ","))
	}

	// Build Event Map
	eventMap := make(map[int]Event)
	for i, record := range records {
		// Skip the first line
		if i == 0 {
			continue
		}
		fmt.Printf("Parsed record: %s\n", record)
		thisId, _ := strconv.Atoi(strings.TrimSpace(record[0]))
		thisName := strings.TrimPrefix(record[1], " ")
		nJump, _ := strconv.Atoi(strings.TrimSpace(record[2]))
		thisJumpList := make([]EventJumpOption, nJump)
		var sumP float32 = 0.0
		for j := 0; j < nJump; j++ {
			thisP64, _ := strconv.ParseFloat(strings.TrimSpace(record[3+2*j]), 32)
			thisP := float32(thisP64)
			sumP += thisP
			thisJumpId, _ := strconv.Atoi(strings.TrimSpace(record[4+2*j]))
			thisJumpList[j] = EventJumpOption{thisP, thisJumpId}
		}
		if sumP > 1.0 {
			fmt.Printf("Error: In Event %d, sum of Prababilities is larger than 1.", thisId)
			return
		}
		eventMap[thisId] = Event{thisId, thisName, thisJumpList}
		printEvent(eventMap[thisId])
	}

	// Simulate through the Event Map
	fmt.Println("Event Simulation Started.")
	currentEventId := 0
	for true {
		fmt.Println("----------------------")
		fmt.Printf("Id: %d\n", eventMap[currentEventId].id)
		fmt.Println("Event: " + eventMap[currentEventId].name)
		p := rand.Float32()
		isToJump := false
		var sumP float32 = 0.0
		i := 0
		for i = 0; i < len(eventMap[currentEventId].jumpList); i++ {
			sumP += eventMap[currentEventId].jumpList[i].p
			if p > sumP {
				continue
			} else {
				isToJump = true
				break
			}
		}
		if isToJump {
			currentEventId = eventMap[currentEventId].jumpList[i].jumpID
		}
		fmt.Printf("P: %.2f\n", p)
		if isToJump {
			fmt.Printf("Jump to: %d\n", currentEventId)
		} else {
			fmt.Println("No Jump.")
			break
		}
	}
	fmt.Println("Event Simulation Finished.")
}

func printEvent(e Event) {
	fmt.Printf("Id: %d\n", e.id)
	fmt.Println("Name:" + e.name)
	fmt.Printf("Number of Jumps: %d\n", len(e.jumpList))
	for i, jump := range e.jumpList {
		fmt.Printf("----Jump ID: %d\n", i)
		fmt.Printf("    Jump p: %.2f\n", jump.p)
		fmt.Printf("    Jump to: %d\n", jump.jumpID)
	}
}
