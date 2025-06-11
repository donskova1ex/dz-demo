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

	wg.Add(2)

	go func() {
		defer wg.Done()
		defer close(numbers)
		for i := 0; i < 10; i++ {
			num := rand.Intn(101)
			numbers <- num
		}
	}()

	go func() {
		defer wg.Done()
		defer close(squares)
		for num := range numbers {
			square := num * num
			squares <- square
		}

	}()

	for result := range squares {
		fmt.Printf("[%d] ", result)
	}

	wg.Wait()
}
