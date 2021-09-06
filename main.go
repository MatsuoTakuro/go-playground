package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		var result int
		time.Sleep(2 * time.Second)
		ch <- result
	}()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Time out")
	case result := <-ch:
		fmt.Println(result)
	}
}
