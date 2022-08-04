// https://qiita.com/taigamikami/items/fc798cdd6a4eaf9a7d5e

package main

import (
	"fmt"
	"time"
)

func f(value string) {
	for i := 0; i < 3; i++ {
		fmt.Println(value)
		time.Sleep(3 * time.Second)
	}
}

func a1() {
	go f("goroutine")
	f("normal")
	fmt.Println("done")
}

func a2() {
	messages := make(chan string)
	go func() { messages <- "Hello" }()

	msg := <-messages
	fmt.Println(msg)
}

func a3() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func a4() {
	// deadlock
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func a5() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func a6() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

func a7() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
}

func a8() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 5; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

func main() {
	// a1()
	// a2()
	// a3()
	// a4()
	// a5()
	// a6()
	// a7()
	a8()
}
