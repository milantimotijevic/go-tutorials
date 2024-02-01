package main

import (
	"fmt"
	"sync"
	"time"
)

type Character struct {
	name               string
	class              string
	subclass           string
	mainAttributeName  string
	mainAttributeValue int
}

func characterWorker(id int, channel chan Character, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for character := range channel {
		fmt.Println("Worker", id, "processing character", character.name)
		time.Sleep(time.Second)
		fmt.Printf("Worker %v: %v is a(n) %v (%v) using %v as their main stat with the value of %v\n",
			id, character.name, character.class, character.subclass, character.mainAttributeName, character.mainAttributeValue)
	}
}

func main() {
	fmt.Println("-- Worker Pool --")

	characters := []Character{
		// don't have to specify that each one is a Character, it's inferred from the array type
		{
			name:               "Lae'zel",
			class:              "Fighter",
			subclass:           "Battlemaster",
			mainAttributeName:  "Strength",
			mainAttributeValue: 16,
		},
		{
			name:               "Astarion",
			class:              "Rogue",
			subclass:           "Thief",
			mainAttributeName:  "Dexterity",
			mainAttributeValue: 16,
		},
		{
			name:               "Gale",
			class:              "Wizard",
			subclass:           "Evocation",
			mainAttributeName:  "Intelligence",
			mainAttributeValue: 16,
		},
		{
			name:               "Shadowheart",
			class:              "Cleric",
			subclass:           "Trickery",
			mainAttributeName:  "Wisdom",
			mainAttributeValue: 16,
		},
		{
			name:               "Halsin",
			class:              "Druid",
			subclass:           "Moon",
			mainAttributeName:  "Wisdom",
			mainAttributeValue: 16,
		},
	}

	characterChannel := make(chan Character)
	// spawn a number of workers
	var waitGroup sync.WaitGroup
	for workerId := 1; workerId <= 2; workerId++ {
		// why not a channel reference?
		waitGroup.Add(1)
		go characterWorker(workerId, characterChannel, &waitGroup)
	}

	// push characters into the channel
	for _, character := range characters {
		characterChannel <- character
	}

	// close the channel to signal workers' range blocks they should no longer wait
	close(characterChannel)

	fmt.Println("Waiting for all workers to finish")
	waitGroup.Wait()
	// NOTE: an alternative to a wait group would be to pass another channel that would receive
	// results, then synchronously read those results in the main routine
	fmt.Println("All workers finished")
}
