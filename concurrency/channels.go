package main

import (
	"fmt"
	"math/rand"
	"time"
)

func channels() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for range 10 {
		fmt.Println(<-c)
	}
	d := fanInSelect(boring("Joe"), boring("Ann"))
	for range 10 {
		fmt.Println(<-d)
	}
	fmt.Println("Exiting now")
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

/*
Fan-In function acts as a multiplexer which combines output of 2 channels received. In this function, the infinite for loop listens for new inputs from input1 and input2 and pushes on the same channel c.
*/
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

/* Same fan-in function with select statement */
func fanInSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case msg1 := <-input1:
				c <- msg1
			case msg2 := <-input2:
				c <- msg2
			}
		}
	}()
	return c
}
