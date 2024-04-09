package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan string)
	wg.Add(3)
	go add(1, 2, &wg, ch)
	go subtract(2, 1, &wg, ch)
	go divide(6, 2, &wg, ch)
go func(){

wg.Wait()

close(ch)
}()
	

	// Read from the channel until it's closed
	for result := range ch {
		fmt.Println(result)
	}
}

func add(a, b int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	ch <- fmt.Sprintf("addition is %d", a+b)
}

func subtract(a, b int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	ch <- fmt.Sprintf("subtraction is %d", a-b)
}

func divide(a, b int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()
	ch <- fmt.Sprintf("division is %v", a/b)
}
