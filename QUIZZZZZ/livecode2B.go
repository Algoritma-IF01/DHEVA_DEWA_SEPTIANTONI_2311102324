package main

import (
    "fmt"
)

func main() {
    var M int
    fmt.Print("Masukan jumlah rombongan: ")
    fmt.Scan(&M)

    for i := 0; i < M; i++ {

        var jumlahMenu, banyakOrang int
        var sisa bool
        fmt.Print("Jumlah menu, Banyaknya orang, Apakah makanan itu tersisa? (true/false): ")
        fmt.Scan(&jumlahMenu, &banyakOrang, &sisa)

        var hargatotal int

        if jumlahMenu > 50 {
            hargatotal = 100000
        } else if jumlahMenu > 3 {
            hargatotal = 10000 + (jumlahMenu-3)*2500
        } else {
            hargatotal = 10000
        }

        if sisa {
            hargatotal *= banyakOrang
        }

        fmt.Printf("Total harga yang harus di bayar rombongan %d : Rp %d\n", i+1, hargatotal)
    }
}
