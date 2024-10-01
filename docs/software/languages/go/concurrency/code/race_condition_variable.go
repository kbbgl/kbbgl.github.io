package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)


func stingyCondition(money *int, condition *sync.Cond) {
	for i := 0; i < 1000000; i++ {
		// Uses Lock/Unlock mutex on the condition variable
		condition.L.Lock()
		*money += 10
		// Signals on the condition variable
		// every time we add to the
		// shared money variable
		condition.Signal()
		condition.L.Unlock()
	}
	fmt.Println("Stingy Done")
}


func spendyCondition(money *int, condition *sync.Cond, min int) {
	for i := 0; i < 1000000; i++ {
		condition.L.Lock()
		// Wait while we don't have enough money
		// releasing mutex and suspending execution
		for *money < min {
			condition.Wait()
		}
		// Returning from Wait() reacquires the mutex
		// and substracts the money once there is enough
		*money -= 50
		if *money < 0 {
			fmt.Println("Money is negative!")
			os.Exit(1)
		}
		condition.L.Unlock()
	}
	fmt.Println("Spendy Done")
}


// We want to create a condition where the spender
// never spends money when they have less than some
// threshold (e.g. 50, acting as the condition variable)
func main() {
	money := 100
	minMoney := 50
	mutex := sync.Mutex{}
	condition := sync.NewCond(&mutex)

	go stingyCondition(&money, condition)
	go spendyCondition(&money, condition, minMoney)
	time.Sleep(2 * time.Second)
	println("Money in bank account: ", money)

}