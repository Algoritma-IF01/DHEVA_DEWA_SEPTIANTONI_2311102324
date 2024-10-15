package main

import "fmt"

func main() {
	for {
		var beratKiri, beratKanan float64
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratKiri, &beratKanan)

		if beratKiri < 0 || beratKanan < 0 || beratKiri+beratKanan > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		if beratKiri-beratKanan >= 9 || beratKanan-beratKiri >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
