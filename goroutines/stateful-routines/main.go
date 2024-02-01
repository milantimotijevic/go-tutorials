package main

import (
	"fmt"
	"sync"
)

type ReadRequest struct {
	readItemId      int
	responseChannel chan int
}

type WriteRequest struct {
	writeItemId     int
	writeItemValue  int
	responseChannel chan bool
}

// channels are already a reference
func stateHolder(state map[int]int, readChannel chan ReadRequest, writeChannel chan WriteRequest) {
	for {
		select {
		case readRequest := <-readChannel:
			readRequest.responseChannel <- state[readRequest.readItemId]
		case writeRequest := <-writeChannel:
			state[writeRequest.writeItemId] = writeRequest.writeItemValue
			writeRequest.responseChannel <- true
		}
	}
}

func stateReader(id int, readChannel chan ReadRequest, waitGroup *sync.WaitGroup, numOfReads int) {
	defer waitGroup.Done()
	for i := 0; i < numOfReads; i++ {
		readItemId := id + 100
		readRequest := ReadRequest{readItemId, make(chan int)}
		//fmt.Println("State Reader", id, "requesting to read value at", readRequest.readItemId)
		readChannel <- readRequest
		<-readRequest.responseChannel //readResponse := <-readRequest.responseChannel
		//fmt.Println("State Reader", id, "received value", readResponse)
	}
}

func stateWriter(id int, writeChannel chan WriteRequest, waitGroup *sync.WaitGroup, numOfWrites int) {
	defer waitGroup.Done()
	for i := 0; i < numOfWrites; i++ {
		writeItemId := id + 100
		writeItemValue := id + 1000
		writeRequest := WriteRequest{writeItemId, writeItemValue, make(chan bool)}
		//fmt.Println("State Writer", id, "requesting to write value", writeItemValue, "at", writeItemId)
		writeChannel <- writeRequest
		<-writeRequest.responseChannel //fmt.Println("State Writer", id, "successfully wrote", writeItemValue, "at", writeItemId, ":", writeResponse)
	}
}

func main() {
	fmt.Println("-- Stateful Goroutines --")

	var state = map[int]int{}
	readChannel := make(chan ReadRequest)
	writeChannel := make(chan WriteRequest)

	go stateHolder(state, readChannel, writeChannel)

	var waitGroup sync.WaitGroup
	operationsToRunCount := 1000
	for i := 1; i <= operationsToRunCount; i++ {
		waitGroup.Add(2)
		go stateWriter(i, writeChannel, &waitGroup, 1000)
		go stateReader(i, readChannel, &waitGroup, 1000)
	}

	waitGroup.Wait()
	fmt.Println("All done. State:", state)
}
