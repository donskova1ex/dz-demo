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
    squareChan := make(chan int, 10)

    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()
        for i := 0; i < 10; i++ {
            num := rand.Intn(101)
            numChan <- num
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


    wg.Wait()

    var results []int
    for square := range squareChan {
        results = append(results, square)
    }

    fmt.Println("Квадраты чисел:", results)
}
