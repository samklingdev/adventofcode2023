package main

import "fmt"

func generator(limit int, out chan<- int) {
	for i := 2; i < limit; i++ {
		out <- i
	}
	close(out)
}

func filter(in <-chan int, out chan<- int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
	close(out)
}

func sieve(limit int) {
	ch := make(chan int)

	go generator(limit, ch)

	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve(100)
}
