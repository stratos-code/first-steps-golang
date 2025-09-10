package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channels in go are a way to communicate between goroutines
// they are a typed conduit through which you can send and
// receive values with the channel operator, <-
// channels can be buffered or unbuffered (synchronous or asynchronous)
// unbuffered channels block the sender until the receiver
// receives the value (and viceversa)
func main() {
	// error for deadlock
	// errorDeadlock()
	// example1()
	// example2()
	// example3()
	realisticExample()
}

// Realistic example of using channels to communicate between goroutines
var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func realisticExample() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}

	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1) // simulate time to check prices
		var chickenPrice = rand.Float32() * 20

		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1) // simulate time to check prices
		var tofuPrice = rand.Float32() * 20

		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	// Wait for the first message from either channel
	// fmt.Printf("\nFound a deal on chicken at %s\n", <-chickenChannel)
	// fmt.Printf("\nFound a deal on tofu at %s\n", <-tofuChannel)

	// Using select to wait on multiple channel operations
	// The select statement blocks until one of its cases can run,
	// then it executes that case. It chooses one at random if multiple are ready.
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nFound a deal on chicken at %s\n", website)
	case website := <-tofuChannel:
		fmt.Printf("\nFound a deal on tofu at %s\n", website)
	}
}

func example3() {
	var c = make(chan int, 5) // buffered channel with capacity of 5
	go process3(c)

	// read values from channel
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1) // simulate processing time
	}
}

func process3(c chan int) {
	defer close(c) // close the channel when the function ends
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Process 3 ended")
}

func example2() {
	var c = make(chan int) // unbuffered channel
	go process2(c)

	// read values from channel
	for i := range c {
		fmt.Println(i)
		time.Sleep(time.Second * 1) // simulate processing time
	}
}

func process2(c chan int) {
	defer close(c) // close the channel when the function ends
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Process 2 ended")
}

func example1() {
	var c = make(chan int)
	go process1(c)
	fmt.Println(<-c)
}

func process1(c chan int) {
	c <- 123
}

// Example of deadlock error using unbuffered channel
func errorDeadlock() {
	var c = make(chan int) // canal sin buffer (0 capacity)
	c <- 1                 // adding value to channel (this line will cause a deadlock)
	var i = <-c            // reading value from channel
	fmt.Println(i)
}
