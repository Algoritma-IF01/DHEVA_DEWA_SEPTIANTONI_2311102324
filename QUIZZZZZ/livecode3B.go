package main

import (
	"fmt"
)

func kelipatan4(sum int) int {
	var angka int
	fmt.Scan(&angka)

	if angka < 0 {
		return sum
	}

	if angka > 0 && angka%4 == 0 {
		sum += angka
	}

	return kelipatan4(sum)
}

func main() {
	totalSum := kelipatan4(0)
	fmt.Printf("Jumlah bilangan positif kelipatan 4: %d\n", totalSum)
}
