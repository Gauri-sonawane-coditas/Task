package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	add(1, 2, ch)
	subtract(2, 1, ch)
	divide(6, 2, ch)

	// Close the channel after all operations are done
	close(ch)

	// Read from the channel until it's closed
	for result := range ch {
		fmt.Println(result)
	}
}

func add(a, b int, ch chan string) {
	ch <- fmt.Sprintf("addition is %d", a+b)
}

func subtract(a, b int, ch chan string) {
	ch <- fmt.Sprintf("subtraction is %d", a-b)
}

func divide(a, b int, ch chan string) {
	ch <- fmt.Sprintf("division is %.2f", float64(a)/float64(b))
}
