// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main3() {

	var pil int

	fmt.Println("Menu")
	fmt.Println("1.Menghitung Luas Segitiga")
	fmt.Println("2.Menghitung Luas Lingkaran")
	fmt.Println("3.Menghitung Luas Persegi")
	fmt.Println("4.EXIT")
	fmt.Println("====================")
	fmt.Print("Silahkan pilih menu: ")
	fmt.Scanf("%d", &pil)
	fmt.Println("====================")

	switch pil {
	case 1:
		fmt.Println("Menghitung Luas Segitiga")
		fmt.Println("________________________")
		var a int
		var t int
		var hasil int

		fmt.Print("Masukkan panjang alas segitiga: ")
		fmt.Scanf("%d", &a)
		fmt.Print("Masukkan panjang tinggi segitiga: ")
		fmt.Scanf("%d", &t)

		hasil = (a * t) / 2
		fmt.Printf("Hasil luas segitiga adalah %d cm^2", hasil)

		break
	case 2:
		fmt.Println("Menghitung Luas Lingkaran")
		fmt.Println("________________________")

		var r float64
		var hasil float64

		fmt.Print("Masukkan panjang jari-jari: ")
		fmt.Scanf("%f", &r)

		hasil = (22 / 7) * r * r
		fmt.Printf("Hasil luas segitiga adalah %f cm^2", hasil)
		break
	case 3:
		fmt.Println("Menghitung Luas Persegi")
		fmt.Println("________________________")

		var s int
		var hasil int

		fmt.Print("Masukkan panjang sisi: ")
		fmt.Scanf("%d", &s)

		hasil = s * s
		fmt.Printf("Hasil luas segitiga adalah %d cm^2", hasil)
		break
	default:
		fmt.Println("====Exit Program====")

	}
}
