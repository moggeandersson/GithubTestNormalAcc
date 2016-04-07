package main

import "fmt"

//We ended up in a deadlock since the channel is recieving and sending at the same time sort of,
//so solve this I separated the recieving part into a separate go routine.
func main() {
    ch := make(chan string)
    go func(){
    	ch <- "Hello world!"
    }()
    fmt.Println(<-ch)
}