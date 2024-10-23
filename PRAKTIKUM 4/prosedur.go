package main

import "fmt"

func fibonacci (n int) int {
	if n == 0 {
        return n
    } else if n == 1 {
		return 1
	}else {
    return fibonacci(n-1) + fibonacci(n-2)
	}
}

func main() {
	fmt.Println("Deret Finobacci hingga deret ke-10: ")
	for i := 0; i <= 10; i++ {
        fmt.Printf("fibonacci(%d) = %d\n", i, fibonacci(i))
    }
}