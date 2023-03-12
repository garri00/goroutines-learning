package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ppcount struct {
	mu sync.Mutex
	v  int
}

func (pp *ppcount) Inc() {
	pp.mu.Lock()

	defer pp.mu.Unlock()
	pp.v++
}

// Do not need this function
func (pp *ppcount) GetValue() int {
	pp.mu.Lock()

	defer pp.mu.Unlock()
	return pp.v
}

func main() {
	pp := ppcount{v: 0}
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Dont figure out how to work with this, but it will help
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				//exit out go routine
				return
			default:
				<-ch2
				pp.Inc()
				fmt.Println("ping : ", pp.GetValue())
				//time.Sleep(time.Second)
				ch1 <- 1
			}

		}

	}()

	go func() {

		for {

			select {
			case <-ctx.Done():
				//exit out go routine
				return
			default:
				ch2 <- 1
				<-ch1
				pp.Inc()
				fmt.Println("pong : ", pp.GetValue())
				//time.Sleep(time.Second)
			}

		}

	}()

	fmt.Println("Timeout")
	time.Sleep(11 * time.Second)
	cancel()

	fmt.Println("PingPong value : ", pp.GetValue())
}
