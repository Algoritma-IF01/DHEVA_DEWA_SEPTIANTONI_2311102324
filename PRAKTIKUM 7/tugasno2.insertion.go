package main

import (
	"bufio"   // Untuk membaca input baris demi baris
	"fmt"     // Untuk mencetak ke layar
	"os"      // Untuk mengambil input dari sistem
	"strconv" // Untuk konversi string ke integer
	"strings" // Untuk memproses string
)

const nMax_2324 = 7919 // Menentukan jumlah maksimum buku yang dapat ditampung

// Struct Buku_2324 untuk menyimpan informasi tentang buku
type Buku_2324 struct {
	id_2324       int    // ID buku
	judul_2324    string // Judul buku
	penulis_2324  string // Penulis buku
	penerbit_2324 string // Penerbit buku
	tahun_2324    int    // Tahun terbit
	rating_2324   int    // Rating buku
}

// Fungsi untuk mengisi data buku ke dalam pustaka
func DaftarkanBuku_2324(pustaka_2324 []Buku_2324, n_2324 int) []Buku_2324 {
	reader_2324 := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input
	for i_2324 := 0; i_2324 < n_2324; i_2324++ {
		fmt.Printf("\nMasukkan data untuk buku ke-%d:\n", i_2324+1)

		// Membaca ID buku
		fmt.Print("ID: ")
		idInput_2324, _ := reader_2324.ReadString('\n')
		idInput_2324 = strings.TrimSpace(idInput_2324)
		id_2324, _ := strconv.Atoi(idInput_2324)

		// Membaca Judul buku
		fmt.Print("Judul: ")
		judul_2324, _ := reader_2324.ReadString('\n')
		judul_2324 = strings.TrimSpace(judul_2324)

		// Membaca Penulis buku
		fmt.Print("Penulis: ")
		penulis_2324, _ := reader_2324.ReadString('\n')
		penulis_2324 = strings.TrimSpace(penulis_2324)

		// Membaca Penerbit buku
		fmt.Print("Penerbit: ")
		penerbit_2324, _ := reader_2324.ReadString('\n')
		penerbit_2324 = strings.TrimSpace(penerbit_2324)

		// Membaca Tahun buku
		fmt.Print("Tahun: ")
		tahunInput_2324, _ := reader_2324.ReadString('\n')
		tahunInput_2324 = strings.TrimSpace(tahunInput_2324)
		tahun_2324, _ := strconv.Atoi(tahunInput_2324)

		// Membaca Rating buku
		fmt.Print("Rating: ")
		ratingInput_2324, _ := reader_2324.ReadString('\n')
		ratingInput_2324 = strings.TrimSpace(ratingInput_2324)
		rating_2324, _ := strconv.Atoi(ratingInput_2324)

		// Menambahkan data ke dalam pustaka
		pustaka_2324[i_2324] = Buku_2324{id_2324, judul_2324, penulis_2324, penerbit_2324, tahun_2324, rating_2324}
	}
	return pustaka_2324
}

// Fungsi untuk mengurutkan buku berdasarkan rating secara menurun menggunakan Insertion Sort
func UrutkanBuku_2324(pustaka_2324 []Buku_2324, n_2324 int) {
	for i_2324 := 1; i_2324 < n_2324; i_2324++ {
		key_2324 := pustaka_2324[i_2324] // Elemen yang akan ditempatkan
		j_2324 := i_2324 - 1

		// Geser elemen yang rating-nya lebih kecil
		for j_2324 >= 0 && pustaka_2324[j_2324].rating_2324 < key_2324.rating_2324 {
			pustaka_2324[j_2324+1] = pustaka_2324[j_2324]
			j_2324--
		}
		pustaka_2324[j_2324+1] = key_2324 // Letakkan elemen di tempat yang sesuai
	}
}

// Fungsi untuk menampilkan buku dengan rating tertinggi
func CetakTerfavorit_2324(pustaka_2324 []Buku_2324, n_2324 int) {
	if n_2324 == 0 { // Jika tidak ada buku
		fmt.Println("Tidak ada buku di perpustakaan.") // Tampilkan pesan
		return
	}

	idxTerfavorit_2324 := 0                             // Indeks buku dengan rating tertinggi
	ratingTertinggi_2324 := pustaka_2324[0].rating_2324 // Ambil rating dari buku pertama

	// Iterasi untuk mencari buku dengan rating tertinggi
	for i_2324 := 1; i_2324 < n_2324; i_2324++ {
		if pustaka_2324[i_2324].rating_2324 > ratingTertinggi_2324 {
			ratingTertinggi_2324 = pustaka_2324[i_2324].rating_2324
			idxTerfavorit_2324 = i_2324
		}
	}

	// Tampilkan buku terfavorit dengan rating tertinggi
	fmt.Println("\nBuku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
		pustaka_2324[idxTerfavorit_2324].judul_2324, pustaka_2324[idxTerfavorit_2324].penulis_2324,
		pustaka_2324[idxTerfavorit_2324].penerbit_2324, pustaka_2324[idxTerfavorit_2324].tahun_2324,
		pustaka_2324[idxTerfavorit_2324].rating_2324)
}

// Fungsi untuk mencetak lima buku dengan rating tertinggi
func CetakSetTerbaru_2324(pustaka_2324 []Buku_2324, n_2324 int) {
	fmt.Println("\nTOP 5 buku dengan rating tertinggi:")
	// Pastikan buku diurutkan sebelum mencetak
	UrutkanBuku_2324(pustaka_2324, n_2324)
	for i_2324 := 0; i_2324 < 5 && i_2324 < n_2324; i_2324++ {
		fmt.Printf("Judul: %s, Rating: %d\n", pustaka_2324[i_2324].judul_2324, pustaka_2324[i_2324].rating_2324)
	}
}

// Fungsi untuk mencari buku berdasarkan rating tertentu
func CariBuku_2324(pustaka_2324 []Buku_2324, n_2324 int, ratingCari_2324 int) {
	ditemukan_2324 := false // Penanda jika buku ditemukan

	// Cari buku dengan rating yang sesuai
	fmt.Printf("\nBuku dengan rating %d:\n", ratingCari_2324)
	for i_2324 := 0; i_2324 < n_2324; i_2324++ {
		if pustaka_2324[i_2324].rating_2324 == ratingCari_2324 {
			// Tampilkan data buku
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n\n",
				pustaka_2324[i_2324].judul_2324, pustaka_2324[i_2324].penulis_2324,
				pustaka_2324[i_2324].penerbit_2324, pustaka_2324[i_2324].tahun_2324,
				pustaka_2324[i_2324].rating_2324)
			ditemukan_2324 = true
		}
	}

	// Jika tidak ada buku yang ditemukan
	if !ditemukan_2324 {
		fmt.Printf("Tidak ada buku dengan rating %d\n", ratingCari_2324)
	}
}

// Program utama
func main() {
	reader_2324 := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input

	// Membaca jumlah buku di perpustakaan
	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	nInput_2324, _ := reader_2324.ReadString('\n')
	nInput_2324 = strings.TrimSpace(nInput_2324)
	nPustaka_2324, _ := strconv.Atoi(nInput_2324)

	// Membuat slice untuk menampung data buku
	var Pustaka_2324 = make([]Buku_2324, nPustaka_2324)

	// Memanggil fungsi untuk mengisi data buku
	Pustaka_2324 = DaftarkanBuku_2324(Pustaka_2324, nPustaka_2324)
	fmt.Println()

	// Menampilkan buku dengan rating tertinggi
	CetakTerfavorit_2324(Pustaka_2324, nPustaka_2324)

	// Menampilkan lima buku dengan rating tertinggi
	CetakSetTerbaru_2324(Pustaka_2324, nPustaka_2324)

	// Mencari buku berdasarkan rating tertentu
	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	ratingCariInput_2324, _ := reader_2324.ReadString('\n')
	ratingCariInput_2324 = strings.TrimSpace(ratingCariInput_2324)
	ratingCari_2324, _ := strconv.Atoi(ratingCariInput_2324)
	CariBuku_2324(Pustaka_2324, nPustaka_2324, ratingCari_2324)
}