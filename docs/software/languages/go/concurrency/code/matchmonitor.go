package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func matchRecorder(events *[]string, mutex *sync.RWMutex) {
	for i := 0; ;i++ {
		mutex.Lock()
		*events = append(*events, "Match event " + strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Appended match event")
	}
}

func clientHandler(events *[]string, mutex *sync.RWMutex, st time.Time) {
	for i := 0; i < 100; i++ {
		mutex.RLock()
		allEvents := copyAllEvents(events)
		mutex.RUnlock()
		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	allEvents := make([]string, 0, len(*matchEvents))
	for _, e := range *matchEvents {
		allEvents = append(allEvents, e)
	}
	return allEvents
}


func main() {
	mutex := sync.RWMutex{}

	numberOfEvents := 10000
	matchEvents := make([]string, 0, numberOfEvents)

	for j := 0; j < numberOfEvents; j++ {
		matchEvents = append(matchEvents, "Match event")
	}

	go matchRecorder(&matchEvents, &mutex)

	start := time.Now()

	for i := 0; i < 5000; i++ {
		go clientHandler(&matchEvents, &mutex, start)
	}
	time.Sleep(100 * time.Second)
}