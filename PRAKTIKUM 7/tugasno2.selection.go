package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

// Fungsi untuk memisahkan bilangan ganjil dan genap dari array_2324
func PisahkanGanjilGenap_2324(array_2324 []int) (ganjil_2324 []int, genap_2324 []int) {
	for _, angka_2324 := range array_2324 {
		if angka_2324%2 == 0 {
			genap_2324 = append(genap_2324, angka_2324) // Tambahkan ke array genap_2324
		} else {
			ganjil_2324 = append(ganjil_2324, angka_2324) // Tambahkan ke array ganjil_2324
		}
	}
	return
}

// Fungsi untuk mengurutkan array_2324 secara ascending (naik) menggunakan Selection Sort
func SelectionSortAsc_2324(array_2324 []int) {
	n_2324 := len(array_2324) // Panjang array
	for i_2324 := 0; i_2324 < n_2324-1; i_2324++ {
		minIndex_2324 := i_2324 // Indeks elemen minimum
		for j_2324 := i_2324 + 1; j_2324 < n_2324; j_2324++ {
			if array_2324[j_2324] < array_2324[minIndex_2324] { // Jika elemen lebih kecil ditemukan
				minIndex_2324 = j_2324 // Perbarui indeks elemen minimum
			}
		}
		// Tukar elemen
		array_2324[i_2324], array_2324[minIndex_2324] = array_2324[minIndex_2324], array_2324[i_2324]
	}
}

// Fungsi untuk mengurutkan array_2324 secara descending (turun) menggunakan Selection Sort
func SelectionSortDesc_2324(array_2324 []int) {
	n_2324 := len(array_2324) // Panjang array
	for i_2324 := 0; i_2324 < n_2324-1; i_2324++ {
		maksIndex_2324 := i_2324 // Indeks elemen maksimum
		for j_2324 := i_2324 + 1; j_2324 < n_2324; j_2324++ {
			if array_2324[j_2324] > array_2324[maksIndex_2324] { // Jika elemen lebih besar ditemukan
				maksIndex_2324 = j_2324 // Perbarui indeks elemen maksimum
			}
		}
		// Tukar elemen
		array_2324[i_2324], array_2324[maksIndex_2324] = array_2324[maksIndex_2324], array_2324[i_2324]
	}
}

func main() {
	scanner_2324 := bufio.NewScanner(os.Stdin) // Membaca input dari pengguna baris demi baris

	fmt.Print("[2324] Masukkan jumlah daerah: ")
	scanner_2324.Scan()
	jumlahDaerah_2324, _ := strconv.Atoi(scanner_2324.Text()) // Konversi jumlah daerah ke integer

	for i_2324 := 0; i_2324 < jumlahDaerah_2324; i_2324++ {
		// Meminta input untuk setiap daerah
		fmt.Printf(">> Masukkan jumlah rumah (angka pertama) dan nomor rumah (pisahkan dengan spasi) untuk daerah ke-%d: ", i_2324+1)
		scanner_2324.Scan()
		baris_2324 := scanner_2324.Text()         // Membaca input pengguna sebagai string
		elemen_2324 := strings.Fields(baris_2324) // Memisahkan string menjadi array string

		// Parsing jumlah rumah dan nomor rumah
		jumlahRumah_2324, _ := strconv.Atoi(elemen_2324[0]) // Konversi jumlah rumah ke integer
		nomorRumah_2324 := make([]int, jumlahRumah_2324)    // Buat array untuk menyimpan nomor rumah
		for j_2324 := 0; j_2324 < jumlahRumah_2324; j_2324++ {
			nomorRumah_2324[j_2324], _ = strconv.Atoi(elemen_2324[j_2324+1]) // Konversi setiap nomor rumah ke integer
		}

		// Pisahkan angka ganjil dan genap
		ganjil_2324, genap_2324 := PisahkanGanjilGenap_2324(nomorRumah_2324)

		// Urutkan ganjil dari kecil ke besar
		SelectionSortAsc_2324(ganjil_2324)

		// Urutkan genap dari besar ke kecil
		SelectionSortDesc_2324(genap_2324)

		// Gabungkan hasil ganjil diikuti genap
		hasil_2324 := append(ganjil_2324, genap_2324...) // Gabungkan ganjil dan genap

		// Tampilkan output untuk daerah tersebut
		fmt.Printf("== Hasil Pengurutan Nomor Rumah daerah ke-%d: ", i_2324+1)
		fmt.Println(strings.Trim(fmt.Sprint(hasil_2324), "[]")) // Cetak hasil tanpa tanda kurung
	}
}