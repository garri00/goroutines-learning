package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	counter := 0

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("gone1")
				close(ch2)
				close(ch1)
				return

			default:
				value := <-ch2
				fmt.Println("ping : ", value)
				ch1 <- value + 1
			}

		}

	}()

	go func() {
		for {

			select {
			case <-ctx.Done():
				fmt.Println("gone2")
				close(ch2)
				close(ch1)
				return

			default:
				ch2 <- counter + 1
				counter = <-ch1
				fmt.Println("pong : ", counter)
			}

		}

	}()

	fmt.Println("Timeout")
	//time.Sleep(5 * time.Second)
	<-ctx.Done()
	cancel()

	defer fmt.Println("PingPong value : ", counter)
}
