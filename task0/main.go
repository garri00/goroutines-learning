package main

import (
	"fmt"
	"sync"
)

var WG = sync.WaitGroup{}
var n = 0

func main() {
	n := [][]int{
		{1, 3, 4},
		{0, 3, 8},
		{4, 7, 7, 35},
	}

	for _, v := range n {
		WG.Add(1)
		go sum(v)
	}
	WG.Wait()
}

func sum(mas []int) {
	defer WG.Done()
	res := 0
	n++
	for _, v := range mas {
		res += v

	}

	fmt.Printf("slice %v : %v \n", n, res)

}
