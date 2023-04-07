package main

import (
	"sync"
	"time"
)

var (
	money = 100
	lock  = sync.Mutex{}
)

func Stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		lock.Unlock()
		time.Sleep((1 * time.Millisecond))
	}

	println("Stingy Done")
}

func Spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money -= 10
		lock.Unlock()
		time.Sleep((1 * time.Millisecond))
	}

	println("Spendy Done")
}

func main() {
	go Spendy()
	go Stingy()
	time.Sleep(3000 * time.Millisecond)
	println(money)
}
