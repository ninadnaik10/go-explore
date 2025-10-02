package main

import (
	"fmt"
	"time"
)

func goroutines() {
	go print()
	fmt.Println("Printing right away")
	time.Sleep(time.Second * 20)
	fmt.Println("Printing after 20 seconds")
	fmt.Println("Exiting now")
}

func print() {
	time.Sleep(time.Second * 10)
	fmt.Println("Printing after 10 seconds")
}
