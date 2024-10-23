// PROGRAM 3 BUAH FUNGSI MATEMATIKA (Fogoh, Gohof, Hofog) 
package main 
 
import ( 
    "fmt" 
) 
 
// Fungsi f mengkuadratkan nilai input 
func f(x_324 int) int { // Mengkuadratkan nilai x_324 
    return x_324 * x_324 
} 
 
// Fungsi g mengurangi 2 dari nilai input 
func g(x_324 int) int { // Mengurangi nilai x_324 dengan 2 
    return x_324 - 2 
} 
 
// Fungsi h menambah 1 ke nilai input 
func h(x_324 int) int { // Menambah nilai x_324 dengan 1 
    return x_324 + 1 
} 
 
// Fungsi fogoh menerapkan f ke hasil dari g yang diterapkan pada h dari x_324 
func fogoh(x_324 int) int { // Menerapkan f(g(h(x_324))) 
    return f(g(h(x_324))) 
} 
 
// Fungsi gohof menerapkan g ke hasil dari h yang diterapkan pada f dari x_324 
func gohof(x_324 int) int { // Menerapkan g(h(f(x_324))) 
    return g(h(f(x_324))) 
} 
 
// Fungsi hofog menerapkan h ke hasil dari f yang diterapkan pada g dari x_324 
func hofog(x_324 int) int { // Menerapkan h(f(g(x_324))) 
    return h(f(g(x_324))) 
} 
 
func main() { 
    var a_324, b_324, c_324 int 
 
    // Meminta pengguna untuk memasukkan nilai a_324, b_324, dan c_324 dalam satu baris dipisahkan oleh spasi 
    fmt.Println("Masukkan nilai a, b, dan c (dipisahkan dengan spasi): ") 
    fmt.Scanf("%d %d %d", &a_324, &b_324, &c_324) // Membaca input pengguna 
 
    // Mencetak hasil dari fogoh, gohof, dan hofog berdasarkan input pengguna 
    fmt.Printf("fogoh(%d) = %d\n", a_324, fogoh(a_324)) // Mencetak hasil dari fogoh(a_324) 
    fmt.Printf("gohof(%d) = %d\n", b_324, gohof(b_324)) // Mencetak hasil dari gohof(b_324) 
    fmt.Printf("hofog(%d) = %d\n", c_324, hofog(c_324)) // Mencetak hasil dari hofog(c_324) 
}