package main

import (
	"fmt"
	"time"
)

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c<- chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// var c chan string = make(chan string)

	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	c1 := make(chan string)
	c2 := make(chan string)

	go func () {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func () {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func () {
		for {
			select {
			case msg1 := <- c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <- c2:
				fmt.Println("Message 2", msg2)
			case <- time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}