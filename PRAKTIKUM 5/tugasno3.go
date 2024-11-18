// PROGRAM UNTUK MEREKAP SKOR PERTANDINGAN 2 BUAH KLUB BOLA YANG BERLAGA
package main

import "fmt" // Mengimpor paket fmt untuk I/O

func main() {
	// Deklarasi variabel untuk nama klub A dan klub B
	var klubA_324, klubB_324 string
	// Deklarasi variabel untuk skor masing-masing klub
	var scoreA_324, scoreB_324 int
	// Slice untuk menyimpan hasil pertandingan
	hasil_324 := []string{}
	// Variabel untuk menghitung jumlah pertandingan
	jumlahMatch_324 := 1

	// Input nama klub A
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA_324)
	// Input nama klub B
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB_324)

	// Loop untuk input skor pertandingan
	for {
		// Input skor pertandingan untuk masing-masing klub
		fmt.Printf("Pertandingan %d: ", jumlahMatch_324)
		fmt.Scan(&scoreA_324, &scoreB_324)

		// Jika skor negatif, keluar dari loop
		if scoreA_324 < 0 || scoreB_324 < 0 {
			break
		}

		// Menentukan pemenang atau draw
		if scoreA_324 > scoreB_324 {
			hasil_324 = append(hasil_324, klubA_324)
		} else if scoreB_324 > scoreA_324 {
			hasil_324 = append(hasil_324, klubB_324)
		} else {
			hasil_324 = append(hasil_324, "Draw")
		}

		// Tambah jumlah pertandingan
		jumlahMatch_324++
	}

	// Menampilkan hasil pertandingan
	fmt.Println("\nDaftar hasil pertandingan:")
	for i_324, poin_324 := range hasil_324 {
		if poin_324 == "Draw" {
			fmt.Printf("Hasil %d: Draw\n", i_324+1)
		} else {
			fmt.Printf("Hasil %d: %s\n", i_324+1, poin_324)
		}
	}

	// Menampilkan pesan bahwa pertandingan selesai
	fmt.Println("Pertandingan selesai")
}