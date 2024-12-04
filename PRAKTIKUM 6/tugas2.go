package main

import (
	"fmt" 
)

const maxKapasitas_2324 = 1000 // Konstanta untuk kapasitas maksimum jumlah ikan

// Fungsi untuk menghitung total berat ikan di setiap wadah dan rata-rata berat per wadah
func hitungBerat_2324(x_2324 int, y_2324 int, beratIkan_2324 []float64) ([]float64, float64) {
	// Hitung jumlah wadah yang diperlukan
	jumlahWadah_2324 := (x_2324 + y_2324 - 1) / y_2324   // Pembulatan ke atas jika ada sisa ikan
	beratWadah_2324 := make([]float64, jumlahWadah_2324) // Slice untuk menyimpan berat total setiap wadah

	// Distribusi berat ikan ke setiap wadah
	for i_2324 := 0; i_2324 < x_2324; i_2324++ { // Iterasi untuk semua ikan
		indeksWadah_2324 := i_2324 / y_2324                         // Tentukan indeks wadah berdasarkan kapasitas
		beratWadah_2324[indeksWadah_2324] += beratIkan_2324[i_2324] // Tambahkan berat ikan ke wadah yang sesuai
	}

	// Hitung rata-rata berat ikan per wadah
	totalBerat_2324 := 0.0 // Variabel untuk menyimpan total berat semua ikan
	for _, berat_2324 := range beratWadah_2324 {
		totalBerat_2324 += berat_2324 // Tambahkan berat wadah ke total
	}
	rataRataBerat_2324 := totalBerat_2324 / float64(jumlahWadah_2324) // Hitung rata-rata berat ikan

	return beratWadah_2324, rataRataBerat_2324 // Kembalikan hasil berupa total berat tiap wadah dan rata-rata berat
}

func main() {
	var x_2324, y_2324 int // Variabel untuk jumlah ikan dan kapasitas wadah

	// Input jumlah ikan (x) dan kapasitas wadah (y)
	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x_2324, &y_2324)

	// Validasi input
	if x_2324 <= 0 || y_2324 <= 0 { // Pastikan jumlah ikan dan kapasitas wadah valid
		fmt.Println("Jumlah ikan dan kapasitas wadah harus lebih dari 0 dan tidak boleh negatif.")
		return // Keluar dari program jika input tidak valid
	}

	if x_2324 > maxKapasitas_2324 { // Validasi maksimum jumlah ikan
		fmt.Printf("Jumlah ikan tidak boleh lebih dari %d.\n", maxKapasitas_2324)
		return // Keluar dari program jika jumlah ikan melebihi kapasitas maksimum
	}

	// Input berat masing-masing ikan
	beratIkan_2324 := make([]float64, x_2324) // Slice untuk menyimpan berat tiap ikan
	fmt.Printf("Masukkan berat ikan sebanyak %d:\n", x_2324)
	for i_2324 := 0; i_2324 < x_2324; i_2324++ { // Iterasi untuk input berat ikan
		fmt.Printf("Berat ikan ke-%d: ", i_2324+1)
		fmt.Scan(&beratIkan_2324[i_2324]) // Baca berat ikan
		if beratIkan_2324[i_2324] < 0 {   // Validasi berat ikan tidak boleh negatif
			fmt.Println("Berat ikan tidak boleh negatif.")
			return // Keluar dari program jika ada berat negatif
		}
	}

	// Panggil fungsi untuk menghitung total berat dan rata-rata
	beratWadah_2324, rataRataBerat_2324 := hitungBerat_2324(x_2324, y_2324, beratIkan_2324)

	// Output total berat ikan di setiap wadah
	fmt.Println("\nTotal berat ikan di setiap wadah adalah:")
	for i_2324, berat_2324 := range beratWadah_2324 {
		fmt.Printf(" Wadah %d: %.2f kg\n", i_2324+1, berat_2324) // Tampilkan berat tiap wadah
	}

	// Output rata-rata berat ikan per wadah
	fmt.Printf("\nBerat rata-rata ikan di setiap wadah adalah: %.2f kg\n", rataRataBerat_2324)
}