// PROGRAM SKENA UNTUK MENCETAK SUKU DERET DARI SUKU AWAL 
package main 

import ( 
   "fmt" // Package untuk format I/O seperti print ke layar 
) 

// cetakDeret mencetak setiap suku dari deret bilangan hingga mencapai 1 sesuai aturan yang diberikan 
func cetakDeret(n_324 int) { // Mencetak deret bilangan mulai dari nilai n_324 
   for n_324 != 1 { 
       fmt.Printf("%d ", n_324) // Mencetak suku saat ini diikuti spasi 
       if n_324%2 == 0 { 
           n_324 /= 2 // Jika n_324 genap, bagi 2 
       } else { 
           n_324 = 3*n_324 + 1 // Jika n_324 ganjil, kalikan 3 dan tambah 1 
       } 
   } 
   fmt.Printf("1\n") // Mencetak 1 sebagai akhir deret 
} 

// main adalah titik masuk utama untuk menjalankan program 
func main() { 
   var n_324 int 
   fmt.Println("Masukkan nilai suku awal (n bilangan positif < 1000000): ") // Meminta pengguna memasukkan nilai awal 
   fmt.Scan(&n_324) 

   if n_324 > 1 && n_324 < 1000000 { // Memastikan nilai n_324 berada dalam batas yang valid 
       cetakDeret(n_324) // Memanggil prosedur untuk mencetak deret 
   } else { 
       fmt.Println("Nilai tidak valid. Harap masukkan nilai antara 1 dan 1000000.") // Menampilkan pesan kesalahan jika n_324 tidak valid 
   } 
}