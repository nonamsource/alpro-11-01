package main

import "fmt"

func main() {
	var r float64

	//Membaca input
	fmt.Scan(&r)

	//Menghitung luas lingkaran
	luas := 3.14 * r * r

	//Menampilkan output
	fmt.Printf("%.2f\n", luas)
}
