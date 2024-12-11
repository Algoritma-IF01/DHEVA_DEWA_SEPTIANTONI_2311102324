package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

// Fungsi untuk mengurutkan array_2324 secara ascending (naik) menggunakan Selection Sort
func SelectionSortAsc_2324(array_2324 []int) {
	n_2324 := len(array_2324) // Panjang array
	for i_2324 := 0; i_2324 < n_2324-1; i_2324++ {
		minIndex_2324 := i_2324 // Indeks elemen minimum
		for j_2324 := i_2324 + 1; j_2324 < n_2324; j_2324++ {
			if array_2324[j_2324] < array_2324[minIndex_2324] { // Bandingkan elemen saat ini dengan elemen minimum
				minIndex_2324 = j_2324 // Perbarui indeks elemen minimum
			}
		}
		// Tukar elemen
		array_2324[i_2324], array_2324[minIndex_2324] = array_2324[minIndex_2324], array_2324[i_2324]
	}
}

func main() {
	scanner_2324 := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari pengguna

	// Meminta jumlah daerah
	fmt.Print("Masukkan jumlah daerah: ")
	scanner_2324.Scan()
	jumlahDaerah_2324, _ := strconv.Atoi(scanner_2324.Text()) // Konversi jumlah daerah ke integer

	for i_2324 := 1; i_2324 <= jumlahDaerah_2324; i_2324++ {
		// Meminta input untuk setiap daerah
		fmt.Printf("Masukkan jumlah rumah (angka pertama) dan nomor rumah (pisahkan dengan spasi) untuk daerah ke-%d: ", i_2324)
		scanner_2324.Scan()
		input_2324 := scanner_2324.Text()        // Membaca input pengguna sebagai string
		parts_2324 := strings.Fields(input_2324) // Memisahkan string menjadi array string

		// Parsing jumlah rumah dan nomor rumah
		jumlahRumah_2324, _ := strconv.Atoi(parts_2324[0]) // Konversi jumlah rumah dari string ke integer
		nomorRumah_2324 := make([]int, jumlahRumah_2324)   // Buat array untuk menyimpan nomor rumah
		for j_2324 := 0; j_2324 < jumlahRumah_2324; j_2324++ {
			nomorRumah_2324[j_2324], _ = strconv.Atoi(parts_2324[j_2324+1]) // Konversi setiap nomor rumah ke integer
		}

		// Urutkan nomor rumah menggunakan Selection Sort (ascending)
		SelectionSortAsc_2324(nomorRumah_2324) // Panggil fungsi untuk mengurutkan nomor rumah

		// Tampilkan hasil untuk daerah saat ini
		fmt.Printf("Hasil Pengurutan Nomor Rumah daerah ke-%d: ", i_2324)
		fmt.Println(strings.Trim(fmt.Sprint(nomorRumah_2324), "[]")) // Cetak nomor rumah terurut tanpa tanda kurung
	}
}