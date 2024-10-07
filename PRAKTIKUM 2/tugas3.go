package main

import (
	"fmt"
	"math"
)

func main() {
	var r_2138 float64
	fmt.Print("jari-jari = ")
	fmt.Scanln(&r_2138)

	volume_2138 := (4.0 / 3.0) * math.Pi * math.Pow(r_2138, 3)

	luas_2138 := 4 * math.Pi * math.Pow(r_2138, 2)

	fmt.Printf("== Bola dengan jari-jari %.0f memiliki volume %.4f dan luas permukaan %.4f\n", r_2138, volume_2138, luas_2138)
}
