package main

import (
	"fmt"
	"sync"
	"time"
)


func countdown(seconds *int, mutex *sync.Mutex) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		*seconds -= 1
		mutex.Unlock()
	}
}

func main() {
	count := 5
	mutex := sync.Mutex{}
	go countdown(&count, &mutex)
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
	}
}
