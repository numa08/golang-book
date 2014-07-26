package main

import (
	"time"
	"fmt"
)

func sleep(d time.Duration) {
	select {
		case <- time.After(d):
		return
	}
}

func main() {
	fmt.Println("sleep")
	sleep(time.Second * 2)
	fmt.Println("awake")
}