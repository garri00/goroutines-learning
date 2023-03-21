package main

import (
	"fmt"
	"sync"
)

func main() {
	mas := []int{2, 5, 6, 1, 0, 4, 3, 7, 9, 8}

	fmt.Println(mas)

	sorted := mergeSort(mas)

	fmt.Println(sorted)
	//var wg sync.WaitGroup
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go mergeSort(mas, &wg)
	//}
	//
	//wg.Wait()
}

func mergeSort(items []int) []int {
	//defer wg.Done()

	if len(items) < 2 {
		return items
	}

	//var wg2 sync.WaitGroup
	//for n := 0; n < 2; n++ {
	//	wg2.Add(1)
	//	go wisper(s, i, n, &wg2)
	//}

	//wg2.Wait()

	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func wisper(s string, i int, n int, wg *sync.WaitGroup) {

	fmt.Printf("some wisper (%v, %v) : %v\n", i, n, s)
	wg.Done()
}
