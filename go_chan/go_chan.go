package main

import (
	"fmt"
	"time"
)

func main() {
	// go fmt.Println("goroutine")
	// fmt.Println("main")

	// for i := 0; i < 3; i++ {
	// 	i := i
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()
	// }
	// time.Sleep(10 * time.Millisecond)

	// ch := make(chan string)

	// go func() {
	// 	ch <- "hi"

	// }()
	// msg := <-ch
	// fmt.Println(msg)

	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		msg := fmt.Sprintf("message #%d", i+1)
	// 		ch <- msg
	// 	}
	// 	close(ch)
	// }()

	// for msg := range ch {
	// 	fmt.Println("got:", msg)
	// }

	// msg = <-ch
	// fmt.Printf("Closed: %#v\n", msg)

	// msg, ok := <-ch
	// fmt.Printf("Closed: %#v (ok=%v)", msg, ok)

	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))
	fmt.Println(sleepSort2(values))

}

func sleepSort(values []int) []int {
	if len(values) == 0 {
		return nil
	}

	ch := make(chan int)
	for _, value := range values {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}(value)
	}

	var sorted []int

	for range values {
		n := <-ch
		sorted = append(sorted, n)
	}

	return sorted
}

func sleepSort2(values []int) []int {
	if len(values) == 0 {
		return nil
	}

	ch := make(chan int)
	for _, value := range values {
		go func(n int) {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}(value)
	}
	close(ch)

	var sorted []int

	for value := range ch {
		sorted = append(sorted, value)
	}

	return sorted
}
