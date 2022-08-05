// https://zenn.dev/hsaki/books/golang-context/viewer/intro

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// chapter02
//
func server() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// chapter03
//
var wg sync.WaitGroup

func generator(done chan struct{}, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()
	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case out <- num:
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func generator2(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num:
			}
		}

		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func a03_1() {
	done := make(chan struct{})
	gen := generator(done, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	close(done)

	wg.Wait()
}

func a03_2() {
	ctx, cancel := context.WithCancel(context.Background())
	gen := generator2(ctx, 2)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	cancel()

	wg.Wait()
}

// chapter04
//
func a04() {
	ctx0 := context.Background()
	ctx1, _ := context.WithCancel(ctx0)

	go func(ctx1 context.Context) {
		ctx2, cancel2 := context.WithCancel(ctx1)

		go func(ctx2 context.Context) {
			go func(ctx2 context.Context) {
				select {
				case <-ctx2.Done():
					fmt.Println("G2-2 canceled")
				}
			}(ctx2)

			select {
			case <-ctx2.Done():
				fmt.Println("G2-1 canceled")
			}
		}(ctx2)

		cancel2()

		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}
	}(ctx1)

	time.Sleep(time.Second)
}

func a04_2() {
	ctx0 := context.Background()
	ctx1, cancel1 := context.WithCancel(ctx0)

	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-1 canceled")
		}
	}(ctx1)

	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-2 canceled")
		}
	}(ctx1)

	cancel1()

	time.Sleep(time.Second)
}

func a04_3() {
	ctx0 := context.Background()

	ctx1, cancel1 := context.WithCancel(ctx0)

	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}
	}(ctx1)

	ctx2, _ := context.WithCancel(ctx0)

	go func(ctx2 context.Context) {
		select {
		case <-ctx2.Done():
			fmt.Println("G2 canceled")
		}
	}(ctx2)

	cancel1()

	time.Sleep(time.Second)
}

func a04_4() {
	ctx0 := context.Background()

	ctx1, _ := context.WithCancel(ctx0)

	go func(ctx1 context.Context) {
		ctx2, cancel2 := context.WithCancel(ctx1)

		go func(ctx2 context.Context) {
			ctx3, _ := context.WithCancel(ctx2)

			go func(ctx3 context.Context) {
				select {
				case <-ctx3.Done():
					fmt.Println("G3 canceled")
				}
			}(ctx3)

			select {
			case <-ctx2.Done():
				fmt.Println("G2 canceled")
			}
		}(ctx2)

		cancel2()

		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}
	}(ctx1)

	time.Sleep(time.Second)
}

func generator3(done chan struct{}, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			}
		}
		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func a05() {
	done := make(chan struct{})
	gen := generator(done, 1)
	deadlineChan := time.After(time.Second)

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		select {
		case result := <-gen:
			fmt.Println(result)
		case <-deadlineChan:
			fmt.Println("timeout")
			break LOOP
		}
	}
	close(done)

	wg.Wait()
}

func generator4(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			}
		}
		close(out)
		fmt.Println("generator closed")
	}()
	return out
}

func a05_2() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	gen := generator4(ctx, 1)

	wg.Add(1)

LOOP:
	for i := 0; i < 5; i++ {
		select {
		case result, ok := <-gen:
			if ok {
				fmt.Println(result)
			} else {
				fmt.Println("timeout")
				break LOOP
			}
		}
	}
	cancel()

	wg.Wait()
}

func generator7(ctx context.Context, num int, userID int, authToken string, traceID int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num:
			}
		}

		close(out)
		fmt.Println("log: ", userID, authToken, traceID)
		fmt.Println("generator closed")
	}()

	return out
}

func a07() {
	ctx, cancel := context.WithCancel(context.Background())
	gen := generator7(ctx, 1, 2, "xxxxxxxx", 3)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	cancel()

	wg.Wait()
}

func generator7_1(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case out <- num:
			}
		}

		close(out)
		userID, authToken, traceID := ctx.Value("userID").(int), ctx.Value("authToken").(string), ctx.Value("traceID").(int)
		fmt.Println("log: ", userID, authToken, traceID)
		fmt.Println("generator closed")
	}()

	return out
}

func a07_1() {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "userID", 2)
	ctx = context.WithValue(ctx, "authToken", "xxxxxxxx")
	ctx = context.WithValue(ctx, "traceID", 3)
	gen := generator7_1(ctx, 1)

	wg.Add(1)

	for i := 0; i < 5; i++ {
		fmt.Println(<-gen)
	}
	cancel()

	wg.Wait()
}

func main() {
	// server()
	// a03_1()
	// a03_2()
	// a04()
	// a04_2()
	// a04_3()
	// a04_4()
	// a05()
	// a05_2()
	a07()
}
