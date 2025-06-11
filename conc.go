package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	sq()
}

func sq() {
	numbers := make(chan int)
	squares := make(chan int)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(numbers)
		for i := 0; i < 10; i++ {
			numbers <- rand.Intn(101)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(squares)
		for num := range numbers {
			squares <- num * num
		}
	}()

	go func() {
		wg.Wait()
	}()

	for result := range squares {
		fmt.Printf("[%d] ", result)
	}

	
}
