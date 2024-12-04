package main  

import (  
	"fmt"
)  

func isValidVoucher(voucher string) bool {  
	panjangdigit := len(voucher)  

	// Kriteria 1: Banyaknya digit harus 5 atau 6  
	if panjangdigit != 5 && panjangdigit != 6 {  
		return false  
	}  

	// Kriteria 2: Perkalian digit dua digit pertama dan dua digit terakhir  
	digitpertama := voucher[:2]  
	digitterakhir := voucher[panjangdigit-2:]  

	firstProduct := (int(digitpertama[0]-'0') * int(digitpertama[1]-'0'))  
	lastProduct := (int(digitterakhir[0]-'0') * int(digitterakhir[1]-'0'))  

	if firstProduct != lastProduct {  
		return false  
	}  

	// Kriteria 3: Digit tengah harus genap  
	var digittengah int  
	if panjangdigit == 5 {  
		digittengah = int(voucher[2] - '0')  
	} else {  
		digittengah = int(voucher[2]-'0')*100 + int(voucher[3]-'0')*10 + int(voucher[4]-'0')  
	}  

	if digittengah%2 != 0 {  
		return false  
	}  

	return true  
}  

func main() {  
	var N int  
	fmt.Println("================================")
	fmt.Print(">Masukkan jumlah voucher mhsw: ")
	fmt.Scan(&N)  

	valid := 0  
	invalid := 0  

	for i := 0; i < N; i++ {  
		var voucher string 
		fmt.Println("================================") 
		fmt.Print(">Masukkan voucher: ")
		fmt.Scan(&voucher)  

		if isValidVoucher(voucher) {  
			fmt.Println("=Voucher valid")
			valid++  
		} else {  
			fmt.Println("=Voucher invalid")
			invalid++  
		}  
	}  
	fmt.Printf("================================\n")
	fmt.Printf("Voucher valid : %d\n", valid)
	fmt.Printf("Voucher invalid : %d\n", invalid)  
}  