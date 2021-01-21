package main

import "fmt"

//Simple example with channels for my docker app
func main() {
	x := make(chan string, 3)

	x <- "Hello"
	fmt.Println(<-x, "First message in channel X")

	x <- "World"
	fmt.Println(<-x, "Second message in channel X")

	x <- "Final test in docker"
	fmt.Println(<-x, "Final message in channel X")
}
