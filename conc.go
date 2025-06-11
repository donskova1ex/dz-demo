package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    numChan := make(chan int)
    squareChan := make(chan int)

    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            numChan <- rand.Intn(101)
        }
        close(numChan)
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        for num := range numChan {
            squareChan <- num * num
        }
        close(squareChan)
    }()

    var results []int
    for square := range squareChan {
        results = append(results, square)
    }

    wg.Wait()

    fmt.Println("Квадраты чисел:", results)
}
