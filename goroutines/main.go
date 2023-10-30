package main

import (
	"fmt"
	"sync"
	"time"
)

type OptionsFunc func(*Options)

type Options struct {
	MaxConnection int
	Id            string
	TLS           bool
}

type Server struct {
	Options
}

func defaultOpts() *Options {
	return &Options{
		MaxConnection: 10,
		Id:            "default",
		TLS:           false,
	}
}

func withTls(opts *Options) {
	opts.TLS = true
}

func withMax(count int) OptionsFunc {
	return func(o *Options) {
		o.MaxConnection = count
	}
}

func NewServer(opts ...OptionsFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(o)
	}
	return &Server{
		Options: *o,
	}
}

func main() {
	scopes()
	// s := NewServer(withTls, withMax(99))
	// fmt.Println(s)
	// // should split odd and even numbers with goroutines synchronously.
	// chn := make(chan int)
	// pingPongChn := make(chan string)
	// wg := sync.WaitGroup{}
	// count := 0
	// wg.Add(4)
	// go printOdd(chn, &wg)
	// go printEven(chn, &wg)
	// go ping(pingPongChn, &wg, &count)
	// go pong(pingPongChn, &wg, &count)
	// wg.Wait()
}

func printOdd(chn chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num, ok := <-chn:
			if !ok {
				return
			}
			fmt.Println("odd", num)
			// time.Sleep(time.Millisecond * 500)
			chn <- num + 1
		}
	}
}

func printEven(chn chan int, wg *sync.WaitGroup) {
	defer func() {
		close(chn)
		wg.Done()
	}()
	chn <- 1
	for {
		select {
		case num, ok := <-chn:
			if !ok || num > 99 {
				fmt.Println("even", num)
				fmt.Println("--------ping pong starts-------")
				return
			}
			fmt.Println("even", num)
			// time.Sleep(time.Millisecond * 1000)
			chn <- num + 1
		}
	}
}

func ping(chn chan string, wg *sync.WaitGroup, count *int) {
	defer wg.Done()

	for {
		select {
		case msg, ok := <-chn:
			if !ok {
				return
			}
			fmt.Println(msg)
			time.Sleep(time.Millisecond * 500)
			chn <- "ping"
		}
	}
}

func pong(chn chan string, wg *sync.WaitGroup, count *int) {
	defer wg.Done()
	chn <- "ping"
	for {
		select {
		case msg := <-chn:
			if *count > 9 {
				close(chn)
				return
			}
			fmt.Println(msg)
			time.Sleep(time.Millisecond * 1000)
			*count++
			chn <- fmt.Sprintf("pong %d", *count)
		}
	}
}

func scopes() {
	x := 10

	{
		x = 20
		fmt.Println(x)
	}
	fmt.Println("x outer", x)
}
