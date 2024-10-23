// PROGRAM GEMA MENCARI PEMENANG DARI DAFTAR PESERTA 
package main 

import ( 
   "bufio"   // Package untuk membaca input dari pengguna melalui terminal 
   "fmt"     // Package untuk format I/O seperti print ke layar 
   "os"      // Package untuk mengakses fungsi sistem operasi, seperti membaca input 
   "strconv" // Package untuk melakukan konversi tipe data, seperti dari string ke integer 
   "strings" // Package untuk memanipulasi string, seperti memisahkan teks dan mengubah huruf besar/kecil 
) 

// hitungSkor menghitung jumlah soal yang berhasil diselesaikan dalam waktu kurang atau sama dengan 300 menit, serta menghitung total waktu yang dibutuhkan. 
func hitungSkor(waktuPengerjaan_324 []int, jumlahSoal_324 *int, totalWaktu_324 *int) { // Menghitung jumlah soal dan total waktu yang valid 
   *jumlahSoal_324 = 0 
   *totalWaktu_324 = 0 
   for _, waktu_324 := range waktuPengerjaan_324 { 
       if waktu_324 <= 300 { // Menghitung hanya soal yang diselesaikan dalam <= 300 menit 
           *jumlahSoal_324++             // Menambah jumlah soal yang berhasil diselesaikan 
           *totalWaktu_324 += waktu_324 // Menambah waktu pengerjaan ke total waktu 
       } 
   } 
} 

// prosesInput memproses input dari pengguna dan menentukan pemenang. 
func prosesInput() { // Memproses input dan menentukan pemenang berdasarkan skor 
   scanner := bufio.NewScanner(os.Stdin) // Membuat scanner untuk membaca input dari pengguna 

   fmt.Println("Masukkan nama dan waktu pengerjaan soal (ketik 'Selesai' untuk mengakhiri input): ") 

   pemenang_324, maxSoal_324, minWaktu_324 := "", -1, 99999 // Inisialisasi pemenang, skor maksimum, dan waktu minimum 

   // Memproses input dari pengguna sampai 'Selesai' diketik 
   for scanner.Scan() { 
       line_324 := scanner.Text() 
       if strings.ToLower(line_324) == "selesai" { // Jika pengguna mengetik 'Selesai', keluar dari loop 
           break 
       } 

       data_324 := strings.Fields(line_324) // Memisahkan input menjadi beberapa bagian (nama dan waktu pengerjaan soal) 
       if len(data_324) != 9 {               // Memeriksa apakah input memiliki 9 bagian (1 nama dan 8 waktu) 
           fmt.Println("Input tidak valid, harap masukkan nama dan 8 waktu pengerjaan soal.") // Memberikan pesan kesalahan jika input tidak valid 
           continue 
       } 

       nama_324 := data_324[0]              // Mengambil nama peserta 
       waktuPengerjaan_324 := make([]int, 8) // Membuat array untuk menyimpan waktu pengerjaan soal 
       for i := 0; i < 8; i++ { 
           waktu_324, error_324 := strconv.Atoi(data_324[i+1]) // Konversi string ke integer 
           if error_324 != nil { 
               fmt.Printf("Input waktu tidak valid: %s\n", data_324[i+1]) // Memberikan pesan kesalahan jika waktu tidak valid 
               continue 
           } 
           waktuPengerjaan_324[i] = waktu_324 // Menyimpan waktu pengerjaan ke array 
       } 

       var jumlahSoal_324, totalWaktu_324 int 
       hitungSkor(waktuPengerjaan_324, &jumlahSoal_324, &totalWaktu_324) // Menghitung jumlah soal yang diselesaikan dan total waktu pengerjaan 

       // Menentukan pemenang berdasarkan jumlah soal yang diselesaikan dan total waktu 
       if jumlahSoal_324 > maxSoal_324 || (jumlahSoal_324 == maxSoal_324 && totalWaktu_324 < minWaktu_324) { 
           pemenang_324, maxSoal_324, minWaktu_324 = nama_324, jumlahSoal_324, totalWaktu_324 // Update pemenang jika memenuhi kriteria 
       } 
   } 

   // Menampilkan hasil kompetisi dan pemenang 
   if pemenang_324 == "" { 
       fmt.Println("Tidak ada peserta yang valid.") // Menampilkan pesan jika tidak ada peserta yang valid 
   } else { 
       fmt.Println("\nPemenang") 
       fmt.Printf("Nama : %s\nJumlah soal yang diselesaikan : %d soal!\nTotal waktu : %d menit!\n", pemenang_324, maxSoal_324, minWaktu_324) // Menampilkan pemenang, jumlah soal, dan total waktu 
   } 
} 

// main adalah titik masuk utama untuk menjalankan program 
func main() { 
   prosesInput() // Memanggil prosedur untuk memproses input dan menentukan pemenang 
}