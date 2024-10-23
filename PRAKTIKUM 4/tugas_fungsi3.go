// PROGRAM FUNGSI UNTUK MENENTUKAN POSISI SEBUAH TITIK SEMBARANG BERADA DI SUATU LINGKARAN ATAU TIDAK 
package main 

import ( 
   "fmt"  // Package untuk format I/O seperti print ke layar 
   "math" // Package untuk perhitungan matematika 
) 

// Fungsi untuk menghitung jarak antara dua titik (x1_324, y1_324) dan (x2_324, y2_324) 
func jarak(x1_324, y1_324, x2_324, y2_324 float64) float64 { 
   // Menggunakan rumus jarak Euclidean 
   return math.Sqrt((x2_324-x1_324)*(x2_324-x1_324) + (y2_324-y1_324)*(y2_324-y1_324)) 
} 

// Fungsi untuk memeriksa apakah titik (x_324, y_324) berada di dalam lingkaran dengan pusat (cx_324, cy_324) dan radius r_324 
func diDalam(cx_324, cy_324, r_324, x_324, y_324 float64) bool { 
   // Mengecek apakah jarak antara titik pusat dan titik sembarang kurang dari atau sama dengan radius 
   return jarak(cx_324, cy_324, x_324, y_324) <= r_324 
} 

func main() { 
   var cx1_324, cy1_324, r1_324 float64 // Variabel untuk pusat dan radius lingkaran 1 
   var cx2_324, cy2_324, r2_324 float64 // Variabel untuk pusat dan radius lingkaran 2 
   var x_324, y_324 float64              // Variabel untuk koordinat titik sembarang 

   // Input untuk lingkaran 1 
   fmt.Print("Masukkan koordinat pusat dan radius lingkaran 1 (cx1 cy1 r1): ") 
   fmt.Scan(&cx1_324, &cy1_324, &r1_324) 

   // Input untuk lingkaran 2 
   fmt.Print("Masukkan koordinat pusat dan radius lingkaran 2 (cx2 cy2 r2): ") 
   fmt.Scan(&cx2_324, &cy2_324, &r2_324) 

   // Input untuk titik sembarang 
   fmt.Print("Masukkan koordinat titik sembarang (x y): ") 
   fmt.Scan(&x_324, &y_324) 

   // Mengecek posisi titik terhadap lingkaran 1 dan lingkaran 2 
   diDalamLingkaran1_324 := diDalam(cx1_324, cy1_324, r1_324, x_324, y_324) // Memeriksa apakah titik berada di dalam lingkaran 1 
   diDalamLingkaran2_324 := diDalam(cx2_324, cy2_324, r2_324, x_324, y_324) // Memeriksa apakah titik berada di dalam lingkaran 2 

   // Menentukan dan menampilkan hasil 
   if diDalamLingkaran1_324 && diDalamLingkaran2_324 { 
       // Jika titik berada di dalam lingkaran 1 dan 2 
       fmt.Println("Titik di dalam lingkaran 1 dan 2") 
   } else if diDalamLingkaran1_324 { 
       // Jika titik hanya berada di dalam lingkaran 1 
       fmt.Println("Titik di dalam lingkaran 1") 
   } else if diDalamLingkaran2_324 { 
       // Jika titik hanya berada di dalam lingkaran 2 
       fmt.Println("Titik di dalam lingkaran 2") 
   } else { 
       // Jika titik berada di luar kedua lingkaran 
       fmt.Println("Titik di luar lingkaran 1 dan 2") 
   } 
}