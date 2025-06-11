package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	numbers := make(chan int)  
	squares := make(chan int) 
	var wg sync.WaitGroup


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
	
	var results []int
	for sq := range squares {
		results = append(results, sq)
	}

	wg.Wait()

	fmt.Println(results)
}
