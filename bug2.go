package main

import "fmt"
import "sync"



// SOLUTION: The problem was that the main function sometimes ended before the Print() could print the last number,
// To solve this I implemented the Waitgroup to make sure that when the for loop in main is done,
//it will wait for the second go routine to finish!
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
    ch := make(chan int)
    go Print(ch, &wg)
    for i := 1; i <= 1100; i++ {
        ch <- i
    }
    close(ch)
    wg.Wait()


}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup) {
    for n := range ch { // reads from channel until it's closed
        fmt.Println(n)
    }
    wg.Done()
}