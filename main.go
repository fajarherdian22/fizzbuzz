package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func fizzBuzz(n int) {
	for i := 1; i < n; i++ {
		if isPrime(i) {
			continue
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FooBar ")
		} else if i%3 == 0 {
			fmt.Print("Foo ")
		} else if i%5 == 0 {
			fmt.Print("Bar ")
		} else {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	fizzBuzz(101)
	time.Sleep(10 * time.Second)
}
