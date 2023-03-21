package main

import (
	"fmt"
	"sync"
)

func main() {
	var word string

	fmt.Scan(&word)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go wisp(word, i, &wg)
	}

	wg.Wait()
}

func wisp(s string, i int, wg *sync.WaitGroup) {

	fmt.Println("wisp (", i, "): ", s)

	var wg2 sync.WaitGroup
	for n := 0; n < 2; n++ {
		wg2.Add(1)
		go wisper(s, i, n, &wg2)
	}

	wg2.Wait()
	wg.Done()
}

func wisper(s string, i int, n int, wg *sync.WaitGroup) {

	fmt.Printf("some wisper (%v, %v) : %v\n", i, n, s)
	wg.Done()
}
