package main

import (
	"fmt"
)

func main() {
	mas := []int{2, 5, 6, 1, 0, 4, 3, 7, 9, 8}
	fmt.Println(mas)

	ch := make(chan []int)
	go mergeSort(mas, ch)

	result := <-ch
	fmt.Println(result)

}

func mergeSort(items []int, result chan []int) {
	if len(items) < 2 {
		result <- items
		return
	}

	middle := len(items) / 2

	ch1 := make(chan []int)
	ch2 := make(chan []int)

	go mergeSort(items[:middle], ch1)
	go mergeSort(items[middle:], ch2)

	first := <-ch1
	second := <-ch2

	result <- merge(first, second)
}

func merge(a []int, b []int) []int {
	var final []int

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
