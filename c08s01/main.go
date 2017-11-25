package main

import (
	"fmt"
)

func main() {
	c := producer(10)
	consumer(c)
}

func producer(l int) chan int {
	out := make(chan int, l)
	for i := 0; i < l; i++ {
		fmt.Println("Writing to channel...")
		out <- i
	}
	close(out)
	return out
}

func consumer(c <-chan int) {

	for v := range c {
		fmt.Printf("Reading value: %v from channel\n", v)
	}

}
