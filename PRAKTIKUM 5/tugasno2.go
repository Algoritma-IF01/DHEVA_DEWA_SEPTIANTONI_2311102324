// PROGRAM OPERASI ARRAY
package main

import (
	"fmt"  // Mengimpor paket fmt untuk I/O
	"math" // Mengimpor paket math untuk operasi matematika
)

// Fungsi untuk menampilkan seluruh elemen array
func tampilkanSEluruhArray(array_324 []int) { // array_324 - array dari tipe int yang ingin ditampilkan
	fmt.Println("Seluruh isi array:", array_324)
}

// Fungsi untuk menampilkan elemen array dengan indeks ganjil
func tampilkanIndeksGanjil(array_324 []int) { // array_324 - array dari tipe int
	fmt.Print("Elemen dengan indeks ganjil: ")
	for _, value_324 := range array_324 { // Loop untuk setiap elemen dalam array
		if value_324%2 != 0 { // Memeriksa apakah elemen adalah ganjil
			fmt.Print(value_324, " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen array dengan indeks genap
func tampilkanIndeksGenap(array_324 []int) { // array_324 - array dari tipe int
	fmt.Print("Elemen dengan indeks genap: ")
	for _, value_324 := range array_324 { // Loop untuk setiap elemen dalam array
		if value_324%2 == 0 && value_324 != 0 { // Memeriksa apakah elemen adalah genap dan bukan nol
			fmt.Print(value_324, " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen array pada indeks kelipatan x_324
func tampilkanKelipatanX(array_324 []int, x_324 int) { // array_324 - array dari tipe int, x_324 - nilai kelipatan indeks
	fmt.Printf("Elemen pada indeks kelipatan %d: ", x_324)
	for i_324 := x_324 - 1; i_324 < len(array_324); i_324 += x_324 { // Loop untuk setiap elemen pada indeks kelipatan x_324
		if array_324[i_324] != 0 { // Memeriksa apakah elemen bukan nol
			fmt.Print(array_324[i_324], " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen array pada indeks tertentu
func hapusElemen(array_324 []int, idx_324 int) []int { // array_324 - array dari tipe int, idx_324 - indeks yang akan dihapus
	return append(array_324[:idx_324], array_324[idx_324+1:]...) // Mengembalikan array baru dengan elemen yang telah dihapus
}

// Fungsi untuk menghitung rata-rata elemen array
func tampilkanRataRata(array_324 []int) float64 { // array_324 - array dari tipe int
	sum_324 := 0                           // Variabel untuk menyimpan jumlah elemen
	for _, value_324 := range array_324 { // Loop untuk setiap elemen dalam array
		sum_324 += value_324 // Menambahkan elemen ke jumlah total
	}
	return float64(sum_324) / float64(len(array_324)) // Mengembalikan rata-rata dalam tipe float64
}

// Fungsi untuk menghitung deviasi standar dari elemen array
func tampilkanStandarDeviasi(array_324 []int, rata float64) float64 { // array_324 - array dari tipe int, rata - rata-rata elemen array
	sumSquares_324 := 0.0                  // Variabel untuk menyimpan jumlah kuadrat selisih
	for _, value_324 := range array_324 { // Loop untuk setiap elemen dalam array
		sumSquares_324 += math.Pow(float64(value_324)-rata, 2) // Menghitung kuadrat selisih setiap elemen dari rata-rata
	}
	return math.Sqrt(sumSquares_324 / float64(len(array_324))) // Mengembalikan akar kuadrat dari jumlah kuadrat selisih dibagi jumlah elemen
}

// Fungsi untuk menghitung frekuensi kemunculan elemen tertentu dalam array
func tampilkanFrekuensi(array_324 []int, target_324 int) int { // array_324 - array dari tipe int, target_324 - nilai yang ingin dihitung frekuensinya
	count_324 := 0                         // Variabel untuk menghitung frekuensi
	for _, value_324 := range array_324 { // Loop untuk setiap elemen dalam array
	 if value_324 == target_324 { // Memeriksa apakah elemen sama dengan target_324
			count_324++ // Menambahkan 1 ke frekuensi jika cocok
		}
	}
	return count_324 // Mengembalikan frekuensi kemunculan target_324
}

func main() {
	var n_324, x_324, hapusIndeks_324, target_324 int

	// Input jumlah elemen array dari pengguna
	fmt.Print("Masukkan jumlah elemen array (n): ")
	fmt.Scan(&n_324) // Membaca jumlah elemen array

	// Membuat array dengan ukuran sesuai input
	dataArray_324 := make([]int, n_324)        // Membuat array dengan panjang n_324
	for i_324 := 0; i_324 < n_324; i_324++ { // Loop untuk mengisi elemen array
		fmt.Printf("Masukkan nilai untuk elemen ke-%d: ", i_324)
		fmt.Scan(&dataArray_324[i_324]) // Membaca nilai elemen dari pengguna
	}

	// Loop untuk menampilkan menu operasi array
	for {
		fmt.Println("1. Tampilkan seluruh isi array")
		fmt.Println("2. Tampilkan elemen dengan nilai ganjil")
		fmt.Println("3. Tampilkan elemen dengan nilai genap")
		fmt.Println("4. Tampilkan elemen pada indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata dari elemen array")
		fmt.Println("7. Tampilkan standar deviasi elemen array")
		fmt.Println("8. Tampilkan frekuensi dari suatu bilangan")
		fmt.Println("9. Keluar")

		// Input pilihan_324 operasi dari pengguna
		var pilihan_324 int
		fmt.Print("Masukkan Nomor Operasi yang dipilih: ")
		fmt.Scan(&pilihan_324) // Membaca pilihan_324 operasi dari pengguna

		// Menjalankan operasi sesuai pilihan_324 pengguna
		switch pilihan_324 {
		case 1:
			tampilkanSEluruhArray(dataArray_324) // Menampilkan seluruh isi array

		case 2:
			tampilkanIndeksGanjil(dataArray_324) // Menampilkan elemen dengan nilai ganjil

		case 3:
			tampilkanIndeksGenap(dataArray_324) // Menampilkan elemen dengan nilai genap

		case 4:
			fmt.Print("Masukkan nilai (x) untuk kelipatan indeks: ")
			fmt.Scan(&x_324)                           // Membaca nilai x_324 dari pengguna
			tampilkanKelipatanX(dataArray_324, x_324) // Menampilkan elemen pada indeks kelipatan x_324

		case 5:
			fmt.Print("Masukkan indeks yang akan dihapus: ")
			fmt.Scan(&hapusIndeks_324)                                          // Membaca indeks yang akan dihapus
			if hapusIndeks_324 >= 0 && hapusIndeks_324 < len(dataArray_324) { // Memeriksa apakah indeks valid
				dataArray_324 = hapusElemen(dataArray_324, hapusIndeks_324) // Menghapus elemen pada indeks tertentu
				tampilkanSEluruhArray(dataArray_324)                          // Menampilkan array setelah elemen dihapus
			} else {
				fmt.Println("Indeks tidak valid!") // Menampilkan pesan jika indeks tidak valid
			}

		case 6:
			mean_324 := tampilkanRataRata(dataArray_324)          // Menghitung rata-rata elemen array
			fmt.Printf("Rata-rata elemen array: %.2f\n", mean_324) // Menampilkan rata-rata elemen array

		case 7:
			mean_324 := tampilkanRataRata(dataArray_324)                            // Menghitung rata-rata elemen array
			standarDeviasi_324 := tampilkanStandarDeviasi(dataArray_324, mean_324) // Menghitung deviasi standar elemen array
			fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi_324)   // Menampilkan deviasi standar elemen array

		case 8:
			fmt.Print("Masukkan nilai untuk mencari frekuensi: ")
			fmt.Scan(&target_324)                                                         // Membaca nilai target_324 dari pengguna
			frekuensi_324 := tampilkanFrekuensi(dataArray_324, target_324)              // Menghitung fre kuensi target_324 dalam array
			fmt.Printf("Frekuensi %d dalam array: %d kali\n", target_324, frekuensi_324) // Menampilkan frekuensi target_324

		case 9:
			fmt.Println("Keluar dari program.") // Keluar dari program
			return

		default:
			fmt.Println("Pilihan tidak valid!") // Menampilkan pesan jika pilihan_324 tidak valid
		}
	}
}