package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// generate 10 random number
	// ForRandom()

	// generate random number, loop if not match
	ForRandomMatch()
}

func ForRandom() {
	// set random seed (go version < 1.20)
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 10; i++ {
		n := rand.Intn(10)
		fmt.Println(n)
	}
}

func ForRandomMatch() {
	rand.Seed(time.Now().UnixNano())
	var n int = 0

	for n != 10 {
		n = rand.Intn(30)
		fmt.Println(n)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println("Finish!!!")
}
