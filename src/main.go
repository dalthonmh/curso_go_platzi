package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Hola")
	wg.Add(1)
	go say("world", &wg)

	wg.Wait()

	go func(text string) {
		fmt.Println("Adios", text)
	}("Dalthon")

	time.Sleep(time.Second * 1)

}
