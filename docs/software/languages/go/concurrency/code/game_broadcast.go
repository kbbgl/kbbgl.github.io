package main

import (
	"fmt"
	"sync"
	"time"
)


func playerHandler(condition *sync.Cond, playersRemaining *int, id int) {
	condition.L.Lock()
	fmt.Println(id, ": Connected")
	*playersRemaining--
	if *playersRemaining == 0 {
		condition.Broadcast()
	}
	for *playersRemaining > 0 {
		fmt.Println(id, ": Waiting for more players")
		condition.Wait()
	}
	condition.L.Unlock()
	fmt.Println("All players connected, Ready player",id)
}


func main(){

	condition := sync.NewCond(&sync.Mutex{})
	numberOfPlayers := 4

	for playedId := 0; playedId < 4; playedId++ {
		go playerHandler(condition, &numberOfPlayers, playedId)
		time.Sleep(1 * time.Second)
	}

}