package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

// Fungsi untuk mengurutkan array_2324 secara ascending (naik) menggunakan Insertion Sort
func InsertionSort_2324(array_2324 []int) {
	n_2324 := len(array_2324) // Panjang array
	for i_2324 := 1; i_2324 < n_2324; i_2324++ {
		key_2324 := array_2324[i_2324]                     // Elemen yang sedang diurutkan
		j_2324 := i_2324 - 1                               // Indeks elemen sebelumnya
		for j_2324 >= 0 && array_2324[j_2324] > key_2324 { // Geser elemen yang lebih besar
			array_2324[j_2324+1] = array_2324[j_2324]
			j_2324--
		}
		array_2324[j_2324+1] = key_2324 // Tempatkan elemen di posisi yang benar
	}
}

// Fungsi untuk memeriksa apakah array_2324 memiliki jarak tetap
func IsJarakTetap_2324(array_2324 []int) (bool_2324 bool, jarak_2324 int) {
	if len(array_2324) < 2 { // Jika elemen kurang dari 2
		return true, 0 // Secara default dianggap jarak tetap
	}
	jarak_2324 = array_2324[1] - array_2324[0] // Hitung jarak awal
	for i_2324 := 2; i_2324 < len(array_2324); i_2324++ {
		if array_2324[i_2324]-array_2324[i_2324-1] != jarak_2324 { // Jika jarak berbeda
			return false, 0 // Tidak berjarak tetap
		}
	}
	return true, jarak_2324 // Jika jarak tetap, kembalikan true dan jarak
}

func main() {
	scanner_2324 := bufio.NewScanner(os.Stdin) // Scanner untuk membaca input

	// Meminta jumlah kasus
	fmt.Print("Masukkan jumlah kasus: ")
	scanner_2324.Scan()
	jumlahKasus_2324, _ := strconv.Atoi(scanner_2324.Text()) // Konversi jumlah kasus ke integer

	for kasus_2324 := 1; kasus_2324 <= jumlahKasus_2324; kasus_2324++ {
		// Membaca input untuk setiap kasus
		fmt.Printf("Masukkan data untuk kasus ke-%d (pisah dengan spasi dan akhiri dengan bilangan negatif): ", kasus_2324)
		scanner_2324.Scan()
		baris_2324 := scanner_2324.Text()         // Membaca input sebagai string
		elemen_2324 := strings.Fields(baris_2324) // Memisahkan string menjadi array
		array_2324 := make([]int, 0)              // Array untuk menyimpan bilangan non-negatif

		for _, el_2324 := range elemen_2324 { // Iterasi setiap elemen input
			num_2324, _ := strconv.Atoi(el_2324) // Konversi string ke integer
			if num_2324 < 0 {                    // Jika negatif, hentikan proses
				break
			}
			array_2324 = append(array_2324, num_2324) // Tambahkan ke array
		}

		// Urutkan array menggunakan Insertion Sort
		InsertionSort_2324(array_2324)

		// Periksa apakah array berjarak tetap
		jarakTetap_2324, jarak_2324 := IsJarakTetap_2324(array_2324)

		// Cetak hasil pengurutan
		fmt.Printf("Hasil Pengurutan: %v\n", array_2324)
		if jarakTetap_2324 {
			fmt.Printf("Data berjarak %d\n\n", jarak_2324) // Jika jarak tetap, cetak jaraknya
		} else {
			fmt.Println("Data berjarak tidak tetap\n") // Jika tidak tetap, cetak pesan ini
		}
	}
}