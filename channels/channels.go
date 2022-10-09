package main

import "fmt"

func addUp(a, b int, ch chan<- int) {
	x := a + b
	ch <- x
}

func main() {
	// make a new channel
	adding := make(chan int)

	go addUp(4, 5, adding)
	go addUp(9, 8, adding)

	// Listening creates a queue, so getting the items will be in reverse order.
	msg2 := <-adding
	msg := <-adding
	fmt.Println(msg)
	fmt.Println(msg2)

}
