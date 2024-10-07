package main
import "fmt"

func main() {
	var a, b, c float64
	var hipotenusa bool
	fmt.Print("Masukan nilai A :")
	fmt.Scanln(&a)
	fmt.Print("Masukan nilai B : ")
	fmt.Scanln(&b)
	fmt.Print("Masukan nilai C : ")
	fmt.Scanln(&c)
	hipotenusa = (c*c) == (a*a + b*b)
	fmt.Println("Sisi c adalah hipotenusa segitiga a, b, c : ", hipotenusa)
}