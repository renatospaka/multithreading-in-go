package main

import "time"

var money = 100

func stingy() {
	for i := 1; i <= 1000; i++ {
		money += 10
		time.Sleep((1 * time.Millisecond))
	}

	println("Stingy Done")
}

func Spendy() {
	for i := 1; i <= 1000; i++ {
		money -= 10
		time.Sleep((1 * time.Millisecond))
	}

	println("Spendy Done")
}

func main() {
	go Spendy()
	go stingy()
	time.Sleep(800 * time.Millisecond)
	println(money)
}